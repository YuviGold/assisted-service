#!/usr/bin/env python3

import json
import logging
import os
from pathlib import Path

import subprocess
import requests

from retry import retry
from utils import check_call, check_output

OCP_VERSIONS_FILE = Path("default_ocp_versions.json")
RHCOS_URL = "https://mirror.openshift.com/pub/openshift-v4/dependencies/rhcos/%s/latest/rhcos-live.x86_64.iso"
OUTPUT_DIR = Path("build") / "rhcos"


def main():
    with OCP_VERSIONS_FILE.open("r") as file_stream:
        ocp_versions = json.load(file_stream)

    os.makedirs(OUTPUT_DIR, exist_ok=True)

    for key, metadata in ocp_versions.items():
        verify_image_version(key, metadata["release_image"])

        try: 
            download_rhcos_image(key)
        except subprocess.CalledProcessError:
            logging.error(f"Failed to download rhcos version {key}")


@retry(exceptions=subprocess.CalledProcessError, tries=3, delay=2)
def download_rhcos_image(key: str):
    url = RHCOS_URL % key
    dest = OUTPUT_DIR / f"rhcos-{key}.iso"

    if dest.exists():
        logging.info(f"Skipping {key}. {dest} already exists")
        return

    logging.info(f"Download rhcos version {key} from {url}")

    check_call(f"curl -fSL {url} -o {dest}")


def verify_image_version(ocp_version: str, release_image: str) -> bool:
    segments = get_oc_version(release_image).split(".")
    assert ocp_version == f"{segments[0]}.{segments[1]}", f"{release_image} image version is {segments[0]}.{segments[1]} not {ocp_version}"


def get_oc_version(release_image: str) -> str:
    return check_output(f"oc adm release info '{release_image}' -o template --template {{{{.metadata.version}}}}")


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, format='%(levelname)-10s %(message)s')
    main()
