#!/bin/python
import sys
from xml.etree import ElementTree


if __name__ == "__main__":
    try:
        version = sys.argv[1]
        filename = sys.argv[2]
        xml_tree = ElementTree.parse(filename)
        root = xml_tree.getroot()
        version_element = root.find("PropertyGroup").find("Version")
        version_element.text = version
        xml_tree.write(filename, encoding="unicode")
        print(f"{filename} set to version {version}")
    except Exception as err:
        print(f"Setting version failed with error: {err}")


