import 'dart:convert';
import 'dart:html';
import 'formstate.dart';
import 'pathlookup.dart';
import 'app.dart';

class LoginForm extends FormState {
  EmailInputElement _email;
  PasswordInputElement _password;

  LoginForm(
      String idElem, String emailElem, String passwordElem, String submitBtn)
      : super(idElem, submitBtn) {
    _email = querySelector(emailElem);
    _password = querySelector(passwordElem);

    querySelector(submitBtn).onClick.listen(onSend);
    registerFormElements([_email, _password]);
  }

  String get email {
    return _email.value;
  }

  String get password {
    return _password.value;
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
    afterSend(obj['Data']);
  }

  void afterSend(String sessionID) {
    var finalURL = window.localStorage['return'];
    finalURL += "?token=" + sessionID;

    window.location.replace(finalURL);
  }
}
