#!/bin/bash

set -ex

docker run -it --rm -v $(pwd)/test_size.sh:/test_size.sh -e PKG golang sh /test_size.sh
