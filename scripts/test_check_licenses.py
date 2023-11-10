import os

import check_licenses as cl
import pytest

_FILE_WITH_LICENSE = "./app/app.go"
_FILE_WITH_GETH_LICENSE = "./x/evm/statedb/journal.go"


@pytest.fixture
def cleanup():
    yield
    file = "test.txt"
    if os.path.exists(file):
        os.remove(file)


def test_check_license_in_file_no_license():
    file = "test.txt"
    with open(file, "w") as f:
        f.write("test")
    assert cl.check_license_in_file(file, cl.ENCL_LICENSE) is False


def test_check_license_in_file_geth_license():
    assert cl.check_license_in_file(_FILE_WITH_GETH_LICENSE, cl.ENCL_LICENSE) == "geth"


def test_check_license_in_file_license():
    assert cl.check_license_in_file(_FILE_WITH_LICENSE, cl.ENCL_LICENSE) is True


def test_check_license_in_file_generated(cleanup):
    file = "test.txt"
    with open(file, "w") as f:
        f.write("// Code generated by go generate; DO NOT EDIT.")
    assert cl.check_license_in_file(file, cl.ENCL_LICENSE) == "generated"


def test_check_if_in_exempt_files_not_included():
    file = "/Users/malte/dev/go/kato114/byte/app/app.go"
    assert cl.check_if_in_exempt_files(file) is False


def test_check_if_in_exempt_files_included():
    file = "/Users/malte/dev/go/kato114/byte/x/revenue/v1/genesis.go"
    assert cl.check_if_in_exempt_files(file) is True

    file = "/Users/malte/dev/go/kato114/byte/x/claims/genesis.go"
    assert cl.check_if_in_exempt_files(file) is True

    file = "/Users/malte/dev/go/kato114/byte/x/erc20/keeper/proposals.go"
    assert cl.check_if_in_exempt_files(file) is True

    file = "/Users/malte/dev/go/kato114/byte/x/erc20/types/utils.go"
    assert cl.check_if_in_exempt_files(file) is True

    file = "/Users/malte/dev/go/kato114/byte/x/evm/genesis.go"
    assert cl.check_if_in_exempt_files(file) is False
