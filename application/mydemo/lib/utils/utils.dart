import 'package:device_info_plus/device_info_plus.dart';
import 'package:shared_preferences/shared_preferences.dart';

Future<String> getDeviceInfo() async {
  // Lấy thông tin thiết bị
  DeviceInfoPlugin deviceInfo = DeviceInfoPlugin();
  return deviceInfo.toString();
  // AndroidDeviceInfo androidInfo = await deviceInfo.androidInfo;
  // print('Running on ${androidInfo.model}');  // e.g. "Moto G (4)"

  // IosDeviceInfo iosInfo = await deviceInfo.iosInfo;
  // print('Running on ${iosInfo.utsname.machine}');  // e.g. "iPod7,1"

  // WebBrowserInfo webBrowserInfo = await deviceInfo.webBrowserInfo;
  // print('Running on ${webBrowserInfo.userAgent}');
}

Future<void> saveUserInfo(String userId, String token) async {
  SharedPreferences prefs = await SharedPreferences.getInstance();
  prefs.setString('userId', userId);
  prefs.setString('token', token);
}

Future<String> getToken() async {
  SharedPreferences prefs = await SharedPreferences.getInstance();
  return prefs.getString('token') ?? '';
}

Future<String> getUserId() async {
  SharedPreferences prefs = await SharedPreferences.getInstance();
  return prefs.getString('userId') ?? 'UserID';
}

Future<bool> userIsAdmin() async {
  final uId = await getUserId();
  return uId == 'admin';
}
