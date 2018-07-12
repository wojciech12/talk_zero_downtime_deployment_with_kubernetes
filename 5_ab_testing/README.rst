##########################################
A/B Testing | bucket | Weigthed Deployment 
##########################################

Istio:

::

  route:
  - tags:
    version: v1.0.0
    weight: 90
  - tags:
    version: v2.0.0
    weight: 10

Traefik (see `docs <https://docs.traefik.io/configuration/backends/kubernetes/>`_):

::

  service_backend1: 12.50%
  service_backend2: 12.50%
  service_backend3: 75 # Same as 75%, the percentage sign is optional

Notice: I would recommend handling any A/B in your application. It is too vital to your business.