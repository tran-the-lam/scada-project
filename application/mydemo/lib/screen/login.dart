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

  bool _isPwdVisible = false;

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
      final deviceInfo = getDeviceInfo();

      print("userID: $userID");
      print("password: $password");
      print("deviceInfo: $deviceInfo");

      // Gọi API để kiểm tra thông tin đăng nhập
      final token = await http.login(userID, password, getDeviceInfo());
      if (token.length > 0) {
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
    return MaterialApp(
      // title: 'Profile',
      debugShowCheckedModeBanner: false,
      home: Scaffold(
        // appBar: AppBar(
        //   title: Text('Login'),
        // ),
        backgroundColor: Color.fromARGB(255, 251, 252, 238),
        body: _isLoading
            ? Center(child: CircularProgressIndicator())
            : Center(
                
                child: Padding(
                  padding: EdgeInsets.all(16.0),
                  child: Container(
                    width: 500,
                    height: 400,
                    child: Column(
                      children: [
                        Text(
                          'SCADA',
                          style: TextStyle(fontSize: 30, fontWeight: FontWeight.bold, fontFamily: 'Lobster'),
                        ),
                        SizedBox(height: 80.0),
                        TextField(
                          controller: _userIDController,
                          decoration: InputDecoration(
                            labelText: 'Mã nhân viên',
                            labelStyle:
                                TextStyle(color: Colors.blue, fontSize: 18),
                            enabledBorder: OutlineInputBorder(
                              borderSide: BorderSide(color: Colors.blue),
                              borderRadius: BorderRadius.circular(10),
                            ),
                            focusedBorder: OutlineInputBorder(
                              borderSide: BorderSide(color: Colors.green),
                              borderRadius: BorderRadius.circular(10),
                            ),
                            prefixIcon:
                                Icon(Icons.perm_identity, color: Colors.blue),
                          ),
                        ),
                        SizedBox(height: 30.0),
                        TextField(
                          controller: _passwordController,
                          obscureText: !_isPwdVisible,
                          decoration: InputDecoration(
                            labelText: 'Mật khẩu',
                            labelStyle:
                                TextStyle(color: Colors.blue, fontSize: 18),
                            enabledBorder: OutlineInputBorder(
                              borderSide: BorderSide(color: Colors.blue),
                              borderRadius: BorderRadius.circular(10),
                            ),
                            focusedBorder: OutlineInputBorder(
                              borderSide: BorderSide(color: Colors.green),
                              borderRadius: BorderRadius.circular(10),
                            ),
                            prefixIcon: Icon(Icons.lock, color: Colors.blue),
                            suffixIcon: IconButton(
                              icon: Icon(
                                _isPwdVisible
                                    ? Icons.visibility
                                    : Icons.visibility_off,
                                color: Colors.blue,
                              ),
                              onPressed: () {
                                setState(() {
                                  _isPwdVisible = !_isPwdVisible;
                                });
                              },
                            ),
                          ),
                        ),
                        SizedBox(height: 40.0),
                        ElevatedButton(
                          child: Text('Đăng nhập'),
                          onPressed: _login,
                          style: ButtonStyle(
                            backgroundColor:
                                MaterialStateProperty.all<Color>(Colors.blue),
                            padding: MaterialStateProperty.all<EdgeInsets>(
                              EdgeInsets.symmetric(
                                  horizontal: 40, vertical: 15),
                            ),
                            textStyle: MaterialStateProperty.all<TextStyle>(
                              TextStyle(fontSize: 18),
                            ),
                            shape: MaterialStateProperty.all<
                                RoundedRectangleBorder>(
                              RoundedRectangleBorder(
                                borderRadius: BorderRadius.circular(10),
                              ),
                            ),
                          ),
                        ),
                      ],
                    ),
                  ),
                ),
              ),
      ),
    );
  }
}
