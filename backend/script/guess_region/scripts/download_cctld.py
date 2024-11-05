import csv
import json

import httpx

GEONAME_URL = "https://download.geonames.org/export/dump/countryInfo.txt"

if __name__ == '__main__':
    response = httpx.get(GEONAME_URL)
    data = response.text
    reader = csv.reader(data.split("\n"), delimiter='\t')
    country_mapping = dict()
    for row in reader:
        if len(row) < 2 or row[0].startswith("#"):
            continue
        country_code = row[0].lower()
        country_cctld = row[9]
        if not country_cctld or country_cctld == "":
            continue
        country_mapping[country_cctld] = country_code

    print(json.dumps(country_mapping, indent=2, ensure_ascii=False))
