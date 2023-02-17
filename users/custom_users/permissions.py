from rest_framework.permissions import BasePermission

from . import models


class IsSeller(BasePermission):

    def has_permission(self, request, view):
        return isinstance(request.user, models.Seller)


class IsCustomer(BasePermission):

    def has_permission(self, request, view):
        return isinstance(request.user, models.Customer)
