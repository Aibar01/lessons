from rest_framework_simplejwt.authentication import JWTAuthentication
from rest_framework_simplejwt.exceptions import InvalidToken, AuthenticationFailed
from rest_framework_simplejwt.settings import api_settings
from django.utils.translation import gettext_lazy as _

from . import models


class CustomJWTAuthentication(JWTAuthentication):

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

    def get_user(self, validated_token):
        """
        Attempts to find and return a user using the given validated token.
        """
        if 'seller_id' in validated_token:
            user_id = validated_token['seller_id']
            self.user_model = models.Seller
        elif 'customer_id' in validated_token:
            user_id = validated_token['customer_id']
            self.user_model = models.Customer
        else:
            raise InvalidToken(_("Token contained no recognizable user identification"))

        try:
            user = self.user_model.objects.get(**{api_settings.USER_ID_FIELD: user_id})
        except self.user_model.DoesNotExist:
            raise AuthenticationFailed(_("Seller not found"), code="user_not_found")

        if not user.is_active:
            raise AuthenticationFailed(_("Seller is inactive"), code="user_inactive")

        return user
