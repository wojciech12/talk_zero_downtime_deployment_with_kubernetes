=================================================
Zero deployment of Micro-services with Kubernetes
=================================================

Talk on deployment strategies with Kubernetes covering kubernetes configuration files and the actual implementation of your service in Golang (see `demo <demo>`_).

You will find demos for recreate, rolling updates, blue-green, and canary deployments.

In future, I will aslo provide running examples for a shadow deployment and feature toggle.

Talk from Pizza&Tech Meetup: https://www.meetup.com/meetup-group-nGBiendv/events/255191675/

Slides:

- in `pdf <slides/index.pdf>`_ (source: `slides/ <slides/>`_)
- LinkedIN slideshare

Demos:

- `Recreate <1_demo_recreate>`_
- `Rolling Updates <2_demo_rolling_updates>`_
- `Blue Green <3_demo_bluegreen>`_
- `Canary <4_demo_canary>`_
- `Micro-service implementation with Golang <demo>`_ based on `missy <https://github.com/microdevs/missy>`_

Work-In-Progess: ab testing / feature toggle and shadow deployments. I will also provide a Python example.

Helpful? Please give a *LIKE* to `my Linkedin post about this talk <https://www.linkedin.com/>`_ or a *STAR* to `this github repo <https://github.com/wojciech12/talk_zero_downtime_deployment_with_kubernetes>`_.

How to build your component
===========================

livenessProbe and readinessProbe
--------------------------------

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

- readinessProbe:

  ::

    readinessProbe:
      exec:
        command:
        - cat
         - /tmp/healthy
      initialDelaySeconds: 5
      periodSeconds: 5

Graceful shutdown
-----------------

- handling SIGTERM
- ``health`` and ``ready`` implementation

Check `demo <demo/>`_.

Looking for a Java implementation, check `Mateusz Dyminski talk <https://github.com/mateuszdyminski/zero>`_.


Related Work
============

- https://github.com/mateuszdyminski/zero
- https://kubernetes.io/blog/2018/04/30/zero-downtime-deployment-kubernetes-jenkins/
- https://container-solutions.com/kubernetes-deployment-strategies/
- https://github.com/ContainerSolutions/k8s-deployment-strategies/tree/master/ab-testing
- http://blog.christianposta.com/deploy/blue-green-deployments-a-b-testing-and-canary-releases/
- https://linkerd.io/
- https://istio.io/
- https://github.com/containous/traefik/blob/master/docs/user-guide/kubernetes.md
- https://docs.traefik.io/configuration/backends/kubernetes/
