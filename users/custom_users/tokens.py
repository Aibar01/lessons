from rest_framework_simplejwt.settings import api_settings
from rest_framework_simplejwt import tokens

from . import models


class Token(tokens.Token):

    @classmethod
    def for_user(cls, user):
        """
        Returns an authorization token for the given user that will be provided
        after authenticating the user's credentials.
        """
        user_id = getattr(user, api_settings.USER_ID_FIELD)
        if not isinstance(user_id, int):
            user_id = str(user_id)

        token = cls()
        match user:
            case models.Seller():
                token_claim = 'seller_id'
            case models.Customer():
                token_claim = 'customer_id'
            case _:
                token_claim = 'user_id'

        token[token_claim] = user_id

        return token


class AccessToken(Token):
    token_type = "access"
    lifetime = api_settings.ACCESS_TOKEN_LIFETIME


class RefreshToken(Token):
    token_type = "refresh"
    lifetime = api_settings.REFRESH_TOKEN_LIFETIME

