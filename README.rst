=================================================
Zero deployment of Micro-Services with Kubernetes
=================================================

Do you need to implement a zero-downtime deployment for your product? Did you consider Kubernetes? Check whether my slides, demos of different deployment strategies, and an service implementation example might help you get started. You can run all demos on minikube.

To see how to prepare your microservice, check `Golang <demo>`_ and `.net core <demo_net>`_ implementations (*Looking for contributions in Py and Rust*).

The demos cover recreate, rolling updates, blue-green, and canary deployment strategies. In future, I will add other deployment approaches - see `TBD <https://github.com/wojciech12/talk_zero_downtime_deployment_with_kubernetes#tbd>`_ below.

Slides:

- `Pdf <slides/index.pdf>`_ (source: `slides/ <slides/>`_)

Demos:

- `Recreate <1_demo_recreate>`_
- `Rolling Updates <2_demo_rolling_updates>`_
- `Blue Green <3_demo_bluegreen>`_
- `Canary <4_demo_canary>`_
- `Canary with Traefik <4_demo_canary_traefik>`_
- `Micro-service implementation with Golang <demo>`_ with `missy <https://github.com/microdevs/missy>`_
- `Micro-service implementation with .net core <demo_net>`_. Contribution from `Paweł Ruciński <https://github.com/meanin>`_.

Helpful? Please give a *LIKE* to `the most recent LinkedIn post about this talk <https://www.linkedin.com/feed/update/urn:li:activity:6521644626707824640>`_ or a *STAR* to `this github repo <https://github.com/wojciech12/talk_zero_downtime_deployment_with_kubernetes>`_.

Questions, Feedback? Let me know at wojciech.barczynski@hypatos.ai.

ps. We are hiring - `github.com/hypatos/jobs <https://github.com/hypatos/jobs>`_

Start
=====

You can use a hosted kubernetes, in the demo I assume we are on minikube:

::

  minikube start
  kubectl config use-context minikube

Why minikube? I want to make it easy for you to start.

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

Check `demo <demo/>`_ and `demo_net <demo_net/>`_.

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
7. Python example.
8. Java example.

History
=======

Previous versions of the talk:
 
- Pizza&Tech meetup in Wroclaw - `LinkedIN Slideshare <https://www.slideshare.net/WojciechBarczyski/zero-deployment-of-microservices-with-kubernetes/>`_
- `DevOps Warsaw Meetup at Ocado <https://www.meetup.com/Wroclaw-DevOps-Meetup/events/255394680/>`_
- Golang Warsaw meetup - `pdf export <https://github.com/wojciech12/talk_zero_downtime_deployment_with_kubernetes/tree/meetup_golang_warsaw_2018/slides_short>`_

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
