import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:device_info_plus/device_info_plus.dart';
import 'package:mydemo/utils/http.dart' as http;
// import 'package:http/http.dart' as http;
import 'dart:convert';

class LoginPage extends StatefulWidget {
  @override
  _LoginPageState createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final _formKey = GlobalKey<FormState>();
  String _userID = '';
  String _password = '';

  void saveUserInfo(String userId, String token) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    prefs.setString('userId', userId);
    prefs.setString('token', token);
  }

  String getDeviceInfo() {
    DeviceInfoPlugin deviceInfo = DeviceInfoPlugin();
    return deviceInfo.toString();
  }
TextEditingController _userIDController = TextEditingController();
  TextEditingController _passwordController = TextEditingController();

  bool _isLoading = false;

  Future<void> _login() async {
    setState(() {
      _isLoading = true;
    });

    try {
      final userID = _userIDController.text;
      final password = _passwordController.text;

      // Gọi API để kiểm tra thông tin đăng nhập
      final token = await http.login(userID, password, getDeviceInfo());

      if (token.length > 0) {
        saveUserInfo(userID, token);
        Navigator.pushReplacementNamed(context, '/home');
      } else {
        showDialog(
          context: context,
          builder: (context) => AlertDialog(
            title: Text('Error'),
            content: Text('Invalid userID or password.'),
            actions: [
              ElevatedButton(
                child: Text('OK'),
                onPressed: () {
                  Navigator.of(context).pop();
                },
              ),
            ],
          ),
        );
      }
    } catch (e) {
      print('Error: $e');
    }

    setState(() {
      _isLoading = false;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Login'),
      ),
      body: _isLoading
          ? Center(child: CircularProgressIndicator())
          : Padding(
              padding: EdgeInsets.all(16.0),
              child: Column(
                children: [
                  TextField(
                    controller: _userIDController,
                    decoration: InputDecoration(labelText: 'UserID'),
                  ),
                  TextField(
                    controller: _passwordController,
                    obscureText: true,
                    decoration: InputDecoration(labelText: 'Password'),
                  ),
                  SizedBox(height: 16.0),
                  ElevatedButton(
                    child: Text('Login'),
                    onPressed: _login,
                  ),
                ],
              ),
            ),
    );
  }
}

