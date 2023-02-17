from rest_framework_simplejwt.settings import api_settings
from rest_framework_simplejwt.tokens import AccessToken

from . import models


class CustomAccessToken(AccessToken):

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

    _token_backend = None
