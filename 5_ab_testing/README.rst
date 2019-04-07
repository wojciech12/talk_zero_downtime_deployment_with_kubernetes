##########################################
A/B Testing | bucket | Weigthed Deployment 
##########################################

[WIP]

Istio:

::

  route:
  - tags:
    version: v1.0.0
    weight: 90
  - tags:
    version: v2.0.0
    weight: 10

Better in App:

- https://github.com/rocket-internet-berlin/RocketBucket
- BaaS - google firebase
