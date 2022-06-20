# ./memrun nginx /bin/date
#>> Mon Jun 20 14:29:46 UTC 2022
#
curl -o /tmp/eicar-fileless https://secure.eicar.org/eicar.com && ./memrun nginx /tmp/eicar-fileless