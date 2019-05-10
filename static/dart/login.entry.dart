import 'dart:html';

import 'loginform.dart';

void main() {
  print("Running Login.Entry");
  window.localStorage['return'] = getParameterByName("return");

  new LoginForm("#frmLogin", "#txtIdentity", "#txtPassword", "#btnLogin");
}

String getParameterByName(String name, [String url]) {
  if (url == null) url = window.location.href;

  var uri = Uri.parse(url);
  return uri.queryParameters[name];
}
