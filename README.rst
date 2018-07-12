=================================================
Zero deployment of Golang Service with Kubernetes
=================================================

Work-In-Progress - a series of talks about rolling new versions to production with Kubernetes.


Event: https://www.meetup.com/meetup-group-nGBiendv/events/255191675/

How to build your component
===========================

livenessProbe and readinessprobe
--------------------------------

Slow start of a component, long recovery times:

- livenessProbe:

  ::

        livenessProbe:
          httpGet:
            path: /model
            port: 8000
            httpHeaders:
              - name: X-Custom-Header
                value: Awesome
          initialDelaySeconds: 600
          periodSeconds: 5
          timeoutSeconds: 18
          successThreshold: 1
          failureThreshold: 3

- readinessprobe:

  ::

    readinessProbe:
      exec:
        command:
        - cat
         - /tmp/healthy
      initialDelaySeconds: 5
      periodSeconds: 5

See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/

Graceful shutdown
-----------------

- on SIGTERM

Rolling updates
===============

::

  spec:
    replicas: 4
    strategy:
      type: RollingUpdate
      rollingUpdate:
        maxUnavailable: 0
        maxSurge: 1


Green/Blue
==========

::

    selector:
      app: document-blue

Canary with Traefik
===================

See:

- https://github.com/containous/traefik/blob/master/docs/user-guide/kubernetes.md
- https://docs.traefik.io/configuration/backends/kubernetes/

Bonus - Istio / Linkerd
=======================

- Canary
- Feature Toggle
- A/B

TBD

Related Work
============

- https://kubernetes.io/blog/2018/04/30/zero-downtime-deployment-kubernetes-jenkins/
- https://container-solutions.com/kubernetes-deployment-strategies/
- https://github.com/ContainerSolutions/k8s-deployment-strategies/tree/master/ab-testing
- http://blog.christianposta.com/deploy/blue-green-deployments-a-b-testing-and-canary-releases/
- https://github.com/mateuszdyminski/zero
- https://linkerd.io/
- https://istio.io/
