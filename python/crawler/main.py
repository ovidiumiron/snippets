import urllib2
import operator

URL = "https://gist.githubusercontent.com/ovidiumiron/6536e02b7be4f9f2d80570a14e48ac5b/raw/c4b66fe2e37ff5600f5827305855d4ce553f8dfd/gistfile1.txt"

response = urllib2.urlopen(URL)
data = response.read()
url_count = {}
for line in data.splitlines(True):
    url_count[line] = url_count.get(line, 0) + 1

most_popular = None
max_count = 0

for url, count in url_count.iteritems():
    if count > max_count:
        max_count = count
        most_popular = url

sorted_url = sorted(url_count.items(), key=operator.itemgetter(1))

print most_popular
print sorted_url[-5:]
