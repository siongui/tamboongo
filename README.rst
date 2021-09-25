=======================================
Solution to GO-TAMBOON ไปทำบุญ
=======================================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/tamboongo?status.svg
   :target: https://godoc.org/github.com/siongui/tamboongo

.. image:: https://github.com/siongui/tamboongo/workflows/ci/badge.svg
    :target: https://github.com/siongui/tamboongo/blob/master/.github/workflows/ci.yml

.. image:: https://goreportcard.com/badge/github.com/siongui/tamboongo
   :target: https://goreportcard.com/report/github.com/siongui/tamboongo

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://github.com/siongui/tamboongo/blob/master/UNLICENSE


Development Environment:

  - `Ubuntu 20.04`_
  - `Go 1.17.1`_


API Access
++++++++++

1. Login to `Test Dashboard <https://dashboard.omise.co/test/dashboard>`_ (Skip
   sign up and create a test account)

2. Use credit card information in test data to create tokens via
   `Token API <https://www.omise.co/tokens-api>`_

3. Make donations by creating charges via
   `Charge API <https://www.omise.co/charges-api>`_


Usage
+++++

Get public and secret key from Omise. In root directory of this project,
run:

.. code-block:: bash

  $ export OmisePublicKey="your_omise_public_key_from_test_dashboard"
  $ export OmiseSecretKey="your_omise_private_key_from_test_dashboard"
  $ make

The following is the output:

.. code-block:: bash

  go fmt *.go
  go fmt commandline/*.go
  cd commandline; go run tamboon.go -rot="../fng.1000.csv.rot128"
  performing donations...
  done.
  total received:		 THB	 2686395103
  successful donated:	 THB	 537625452
  faulty donation:	 THB	 2148769651

  average per person:	 THB	 2814792
  top donors:
  	Mrs. Mimosa R Tûk (THB 5075024)
  	Mr. Falco S Bracegirdle (THB 5074457)
  	Mrs. Pimpernel C Headstrong (THB 5068438)


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `GO-TAMBOON ไปทำบุญ <https://github.com/omise/challenges/tree/challenge-go>`_


.. _Go: https://golang.org/
.. _Ubuntu 20.04: https://releases.ubuntu.com/20.04/
.. _Go 1.17.1: https://golang.org/dl/
.. _UNLICENSE: https://unlicense.org/
