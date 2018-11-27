=================================================
Zero deployment of Micro-services with Kubernetes
=================================================

Do you need to implement a zero-downtime deployment for your product? Did you consider Kubernetes? Check whether my slides, demos of different deployment strategies, and an example of a micro-service can help you get started. The demos, you can run easily on minikube. The implementation, currently in `Golang <demo>`_ and `.net core <demo_net>`_ shows how to prepare your service to not lose requests on upgrade.

The demos cover recreate, rolling updates, blue-green, and canary deployment strategies. In future, I will add other deployment approaches - see `TBD <https://github.com/wojciech12/talk_zero_downtime_deployment_with_kubernetes#tbd>`_ below.

So far, I presented this material at: `Pizza&Tech meetup in Wroclaw <https://www.meetup.com/meetup-group-nGBiendv/events/255191675/>`_ and -- the shorter version -- on `Golang Warsaw Meetup <https://www.meetup.com/Golang-Warsaw/events/255260613/>`_.

Slides for 1.5h:

- `LinkedIN Slideshare <https://www.slideshare.net/WojciechBarczyski/zero-deployment-of-microservices-with-kubernetes/>`_
- `Pdf <slides/index.pdf>`_ (source: `slides/ <slides/>`_)

Slides for 40 minutes:

- `Pdf <slides_short/index.pdf>`_ (source: `slides_short/ <slides_short/>`_)

Demos:

- `Recreate <1_demo_recreate>`_
- `Rolling Updates <2_demo_rolling_updates>`_
- `Blue Green <3_demo_bluegreen>`_
- `Canary <4_demo_canary>`_
- `Micro-service implementation with Golang <demo>`_ based on `missy <https://github.com/microdevs/missy>`_

Helpful? Please give a *LIKE* to `the most recent LinkedIn post about this talk <https://www.linkedin.com/feed/update/urn:li:activity:6463041131910352896>`_ or a *STAR* to `this github repo <https://github.com/wojciech12/talk_zero_downtime_deployment_with_kubernetes>`_.

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

TBD
---

To Be Added before the next talk:

1. cover: what happening is during a k8s upgrade
2. cover: `anti-affinity <https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity>`_
3. conver ``pod's disruption budget``
4. shadow deployment with (most probably) Istio
5. feature switch with Golang
6. A/B deployment with Golang
7. Weight-based routing of requests with Traefik
8. Python example.

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
