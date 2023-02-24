from django.contrib.auth import authenticate
from rest_framework import serializers
from django.utils.translation import gettext_lazy as _
from rest_framework_simplejwt import exceptions
from rest_framework_simplejwt.settings import api_settings

from custom_users import tokens
from custom_users.backend import CustomModelBackend


class TokenObtainSerializer(serializers.Serializer):
    phone_number = serializers.CharField()

    default_error_messages = {
        "no_active_account": _("No active account found with the given credentials")
    }

    def validate(self, attrs):
        authenticate_kwargs = {"phone_number": attrs["phone_number"],
                               "request": self.context["request"]}
        self.user = CustomModelBackend().authenticate(**authenticate_kwargs)

        if not api_settings.USER_AUTHENTICATION_RULE(self.user):
            raise exceptions.AuthenticationFailed(
                self.error_messages["no_active_account"],
                "no_active_account",
            )

        return {}

    @classmethod
    def get_token(cls, user):
        return cls.token_class.for_user(user)


class TokenObtainPairSerializer(TokenObtainSerializer):

    def validate(self, attrs):
        data = super().validate(attrs)

        refresh = tokens.RefreshToken.for_user(self.user)
        access = tokens.AccessToken.for_user(self.user)

        data["refresh"] = str(refresh)
        data["access"] = str(access)

        return data
