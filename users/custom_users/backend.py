from django.contrib.auth.backends import ModelBackend
from . import models


class CustomModelBackend(ModelBackend):
    """
    Authenticates against settings.AUTH_USER_MODEL.
    """

    def authenticate(self, request, phone_number=None, **kwargs):
        try:
            user = models.Customer.objects.get(phone_number=phone_number)
        except models.Customer.DoesNotExist:
            return

        if self.user_can_authenticate(user):
            return user
