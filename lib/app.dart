import 'dart:async';
import 'dart:convert';
import 'dart:html';

Future<Map<String, String>> getApp() async {
  identifyLocation();

  var appUrl = window.localStorage['return'];
  var ip = await getIP();
  var location = window.localStorage['location'];
  HiddenInputElement instanceElem = querySelector("#InstanceID");

  return {
    "Name": appUrl,
    "IP": ip,
    "Location": location,
    "InstanceID": instanceElem.value
  };
}

void identifyLocation() {
  if (window.navigator != null) {
    window.navigator.geolocation
        .getCurrentPosition(enableHighAccuracy: true)
        .then((Geoposition position) {
      var location =
          '${position.coords.latitude}, ${position.coords.longitude}';
      window.localStorage['location'] = location;
    }).catchError((err) {
      print('Position Error: ${err.error}');
    });
  }
}

Future<String> getIP() async {
  var resp = await HttpRequest.getString('https://jsonip.com');

  return jsonDecode(resp)["ip"];
}
