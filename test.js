fetch("https://efts.sec.gov/LATEST/search-index", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-US,en;q=0.9,fr;q=0.8",
    "cache-control": "no-cache",
    "content-type": "application/x-www-form-urlencoded; charset=UTF-8",
    "pragma": "no-cache",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-site",
    "Referer": "https://www.sec.gov/",
    "Referrer-Policy": "strict-origin-when-cross-origin"
  },
  "body": "{\"keysTyped\":\"zim\",\"narrow\":true}",
  "method": "POST"
}).then(response => response.json()).then(data => console.log(data));