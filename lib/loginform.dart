import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_entity/bodies/login.dart';
import 'package:mango_entity/secureapi.dart';
import 'package:mango_ui/formstate.dart';

class LoginForm extends FormState {
  EmailInputElement _username;
  PasswordInputElement _password;
  HiddenInputElement _client;
  HiddenInputElement _state;

  LoginForm(String idElem, String clientElem, String stateElem,
      String usernameElem, String passwordElem, String btnSubmit)
      : super(idElem, btnSubmit) {
    _username = querySelector(usernameElem);
    _password = querySelector(passwordElem);
    _client = querySelector(clientElem);
    _state = querySelector(stateElem);

    querySelector(btnSubmit).onClick.listen(onSend);
  }

  String get username {
    return _username.value;
  }

  String get password {
    return _password.value;
  }

  String get client {
    return _client.value;
  }

  String get state {
    return _state.value;
  }

  void onSend(Event e) async {
    if (isFormValid()) {
      disableSubmit(true);

      var req = await sendLogin(new Login(client, username, password));
      var obj = req.response.toString();

      if (req.status == 200) {
        window.location.href =
            "/consent?client=${client}&state=${state}&partial=${obj}";
      } else {
        new Toast.error(
            title: "Error!",
            message: req.response,
            position: ToastPos.bottomLeft);
      }
    }
  }
}
