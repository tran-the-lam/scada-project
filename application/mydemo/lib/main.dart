import 'package:flutter/material.dart';
import 'package:mydemo/screen/login.dart';
import 'package:mydemo/screen/monitor.dart';
import 'package:mydemo/screen/accounts.dart';
import 'package:mydemo/screen/home.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const ProfilePage(title: 'Profile'),
      routes: {
        // '/monitor':(context) => const MonitorPage(title: 'Monitor'),
        // '/users':(context) => const AccountsPage(title: 'Accounts'),
        '/home':(context) => const ProfilePage(title: 'Profile'),
      },
    );
  }
}
