var Items = [];

var i = 0;
var prev = document.querySelector('#prev')
prev.onclick = function() {
  if (i > 0) {
    i--;
  }
  refresh(i, isCom);
  return false;
}

var next = document.querySelector('#next')
next.onclick = function() {
  i++;
  if (i == Items.length) {
    alert("No more!");
    i--;
  }
  refresh(i, isCom);
  return false;
}

var isCom = false;

var com = document.querySelector('#com')
com.onclick = function() {
  if (isCom == true) {
    isCom = false;
  } else {
    isCom = true;
  }
  refresh(i, isCom);
}


function refresh(i, is) {
  if (is == true) {
    link = Items[i]['comments'];
  } else {
    link = Items[i]['link'];
  }
  document.querySelector('#frame').innerHTML = '<iframe src="'+link+'" frameborder="0" nwdisable nwfaketop></iframe>';
  document.querySelector('#link').innerHTML = link;
  
}

var FeedParser = require('feedparser')
  , request = require('request');

var req = request('http://feeds.feedburner.com/newsyc100')
  , feedparser = new FeedParser();

req.on('error', function (error) {
  // handle any request errors
});
req.on('response', function (res) {
  var stream = this;

  if (res.statusCode != 200) {
    return this.emit('error', new Error('Bad status code'));
  }

  stream.pipe(feedparser);
});

feedparser.on('error', function(error) {
  // always handle errors
});
feedparser.on('readable', function() {
  var stream = this
    , item;

  while (item = stream.read()) {
    Items.push(item);
  }
});
