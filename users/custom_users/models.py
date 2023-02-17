import uuid

from django.db import models
from django.contrib.auth.models import AbstractUser, Group, Permission
from django.utils.translation import gettext_lazy as _


class Seller(AbstractUser):
    # id = models.UUIDField(primary_key=True, default=uuid.uuid4)
    company_name = models.CharField(max_length=100, unique=True)
    groups = models.ManyToManyField(
        Group,
        verbose_name=_("groups"),
        blank=True,
        related_name="sellers",
    )
    user_permissions = models.ManyToManyField(
        Permission,
        verbose_name=_("user permissions"),
        blank=True,
        related_name="sellers",
    )

    class Meta:
        verbose_name = _("seller")
        verbose_name_plural = _("sellers")


class Customer(AbstractUser):
    # id = models.UUIDField(primary_key=True, default=uuid.uuid4)
    phone_number = models.CharField(max_length=30)
    groups = models.ManyToManyField(
        Group,
        verbose_name=_("groups"),
        blank=True,
        related_name="customers",
    )
    user_permissions = models.ManyToManyField(
        Permission,
        verbose_name=_("user permissions"),
        blank=True,
        related_name="customers",
    )

    class Meta:
        verbose_name = _("customer")
        verbose_name_plural = _("customers")
