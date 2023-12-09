import 'package:flutter/material.dart';
import 'package:mydemo/screen/monitor.dart';
import 'package:mydemo/screen/accounts.dart';
import 'package:mydemo/screen/profile.dart';
import 'package:mydemo/utils/utils.dart' as utils;

class HomePage extends StatefulWidget {
  const HomePage({super.key, required this.title});
  final String title;

  @override
  State<HomePage> createState() => _BottomNavigationBarExampleState();
}

class _BottomNavigationBarExampleState
    extends State<HomePage> {
  bool isAdmin = false;
  int _selectedIndex = 0;
  static const TextStyle optionStyle = TextStyle(fontSize: 30, fontWeight: FontWeight.bold);

  @override
  void initState() {
    super.initState();
    checkAdmin();
  }

  Future<void> checkAdmin() async {
    final isAdmin = await utils.userIsAdmin();
    setState(() {
      this.isAdmin = isAdmin;
    });
  }

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
      // switch(index){
      //   case 0:
      //     Navigator.pushNamed(context, '/monitor');
      //     break;
      //   case 1:
      //     Navigator.pushNamed(context, '/users');
      //     break;
      //   case 2:
      //     Navigator.pushNamed(context, '/profile');
      //     break;
      // }
    });
  }

  @override
  Widget build(BuildContext context) {

    List<BottomNavigationBarItem> adminItems = [
      BottomNavigationBarItem(
        icon: Icon(Icons.monitor_heart),
        label: 'Monitor',
      ),
      BottomNavigationBarItem(
        icon: Icon(Icons.manage_accounts),
        label: 'Users',
      ),
      BottomNavigationBarItem(
        icon: Icon(Icons.person),
        label: 'Profile',
      ),
    ];
      

    List<BottomNavigationBarItem> userItems = [
      BottomNavigationBarItem(
        icon: Icon(Icons.monitor_heart),
        label: 'Monitor',
      ),
      BottomNavigationBarItem(
        icon: Icon(Icons.person),
        label: 'Profile',
      ),
    ];

    List<BottomNavigationBarItem> items =  isAdmin  ? adminItems : userItems;

    List<Widget> _adminWidgetOptions = <Widget>[
      MonitorPage(title: "Monitor"),
      AccountPage(title: "Accounts"),
      ProfilePage(title: "Profile"),
    ];

    List<Widget> _userWidgetOptions = <Widget>[
      MonitorPage(title: "Monitor"),
      ProfilePage(title: "Profile"),
    ]; 

    List<Widget> _widgetOptions = isAdmin ? _adminWidgetOptions : _userWidgetOptions;
    
    return Scaffold(
      // appBar: AppBar(
      //   title: const Text('BottomNavigationBar Sample'),
      // ),
      body: Center(
        child: _widgetOptions.elementAt(_selectedIndex),
      ),
      bottomNavigationBar: BottomNavigationBar(
        items: items,
        currentIndex: _selectedIndex,
        selectedItemColor: Colors.amber[800],
        onTap: _onItemTapped,
      ),
    );
  }
}