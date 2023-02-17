from rest_framework.decorators import api_view, permission_classes
from rest_framework.response import Response
from rest_framework.views import APIView

from . import permissions


class Index(APIView):
    permission_classes = (permissions.IsSeller,)

    def get(self, request):
        print(type(request.user))
        return Response({'message': 'OK'})