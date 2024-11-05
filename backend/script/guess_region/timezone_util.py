from datetime import datetime

import pytz

# { +0800: ['XXX/XXXXXXXX', ...], ...}
TIMEZONE_OFFSETS = dict()
# { 'America/Los_Angeles': ['us', ...], ...}
TZ_COUNTRIES = dict()


def init_timezone_offsets():
    global TIMEZONE_OFFSETS
    for tz in pytz.all_timezones:
        timezone = pytz.timezone(tz)
        local_time = datetime.now(timezone)
        utc_offset = local_time.utcoffset()
        assert utc_offset is not None
        minutes, remainder = divmod(utc_offset.total_seconds(), 60)
        assert remainder == 0
        hours, minutes = divmod(minutes, 60)
        hours_abs = abs(hours)
        sign = "-" if hours < 0 else "+"
        hours_formatted = sign + str(int(hours_abs)).zfill(2) + str(int(minutes)).zfill(2)
        assert len(hours_formatted) == 5
        if hours_formatted not in TIMEZONE_OFFSETS:
            TIMEZONE_OFFSETS[hours_formatted] = list()
        TIMEZONE_OFFSETS[hours_formatted].append(tz)


def init_tz_counties():
    global TZ_COUNTRIES
    for country, timezones in pytz.country_timezones.items():
        for tz in timezones:
            if tz not in TZ_COUNTRIES:
                TZ_COUNTRIES[tz] = list()
            TZ_COUNTRIES[tz].append(country.lower())


init_timezone_offsets()
init_tz_counties()

def get_countries_by_timezone(tz: str):
    timezones = TIMEZONE_OFFSETS[tz]
    countries = set()
    for timezone in timezones:
        if timezone in TZ_COUNTRIES:
            countries.update(TZ_COUNTRIES[timezone])
    return countries

if __name__ == '__main__':
    print(f"TIMEZONE_OFFSETS: {TIMEZONE_OFFSETS}")
    print(f"TZ_COUNTRIES: {TZ_COUNTRIES}")
