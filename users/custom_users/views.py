from rest_framework.response import Response
from rest_framework.views import APIView
from rest_framework_simplejwt.views import TokenViewBase

from . import permissions, serializers


class TokenObtainPairView(TokenViewBase):
    """
    Takes a set of user credentials and returns an access and refresh JSON web
    token pair to prove the authentication of those credentials.
    """

    serializer_class = serializers.TokenObtainPairSerializer


class Index(APIView):
    permission_classes = (permissions.IsCustomer,)

    def get(self, request):
        print(type(request.user))
        return Response({'message': 'OK'})
