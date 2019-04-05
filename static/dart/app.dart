import 'dart:convert';
import 'dart:html';

Map<String, String> getApp() {
  identifyLocation();

  var appUrl = window.localStorage['return'];
  var ip = window.localStorage['ip'];
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
  if (window.navigator.geolocation != null) {
    window.navigator.geolocation
        .getCurrentPosition(enableHighAccuracy: true)
        .then(setLocation);
  }
}

void setLocation(Geoposition position) {
  var location = position.coords.latitude.toString() +
      ", " +
      position.coords.longitude.toString();
  window.localStorage['location'] = location;
}

Future<String> getIP() async {
  var resp = await HttpRequest.getString('http://jsonip.com');

  return jsonDecode(resp)["ip"];
}
