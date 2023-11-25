Future<void> initAsync() async {
    try {
      await NotificationServices.initializeService();
      //keep same
    } catch (e) {
      log(e.toString());
      navigateToLogin();
    }
  }