import 'dart:convert';
import 'dart:html';
import 'formstate.dart';
import 'pathlookup.dart';

class LoginForm extends FormState {
  EmailInputElement _email;
  PasswordInputElement _password;

  LoginForm(
      String idElem, String emailElem, String passwordElem, String submitBtn)
      : super(idElem, submitBtn) {
    _email = querySelector(emailElem);
    _password = querySelector(passwordElem);

    querySelector(submitBtn).onClick.listen(onSend);
    registerValidation();
  }

  String get email {
    return _email.value;
  }

  String get password {
    return _password.value;
  }

  void registerValidation() {
    _email.onBlur.listen((e) => {validate(e, _email)});
    _password.onBlur.listen((e) => {validate(e, _password)});
  }

  void validate(Event e, InputElement elem) {
    var elemValid = elem.checkValidity();

    if (!elemValid) {
      elem.setAttribute("invalid", "");
    } else {
      elem.removeAttribute("invalid");
    }

    elem.nextElementSibling.text = elem.validationMessage;

    super.disableSubmit(!super.isFormValid());
  }

  void validateArea(Event e, TextAreaElement elem) {
    var elemValid = elem.checkValidity();

    if (!elemValid) {
      elem.setAttribute("invalid", "");
    } else {
      elem.removeAttribute("invalid");
    }

    elem.nextElementSibling.text = elem.validationMessage;

    super.disableSubmit(!super.isFormValid());
  }

  void onSend(Event e) {
    if (isFormValid()) {
      disableSubmit(true);
      submitSend().then((obj) => {disableSubmit(false)});
    }
  }

  Future submitSend() async {
    var url = await buildPath("Secure.API", "login", new List<String>());
    var data =
        jsonEncode({"App": getApp(), "Email": email, "Password": password});

    var req = await HttpRequest.request(url, method: "POST", sendData: data);

    var obj = jsonDecode(req.response);

    print(obj['Data']);
    afterLogin(obj['Data']);
  }

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

  void afterLogin(String sessionID) {
    var finalURL = window.localStorage['return'];
    finalURL += "?token=" + sessionID;

    window.location.replace(finalURL);
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
}
