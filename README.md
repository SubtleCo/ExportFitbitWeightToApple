# ExportFitbitWeightToApple
A go program that converts your Fitbit data dump into a CSV for Apple Health upload

1. [Download your data from Fitbit](https://help.fitbit.com/articles/en_US/Help_article/1133)
2. With go installed, run the script with the directory containing your "weight-*.json" files as the first and only argument, i.e. `go run . ~/Downloads/Takeout/Fitbit/Global_Export_Data` (Tab completion is your friend!)
3. Upload the `export.csv` to somewhere in your iCloud drive
4. Using an iOS device, run [this shortcut](https://www.reddit.com/r/shortcuts/comments/9jszwn/import_weight_data_into_health_app_from_csv/)https://www.reddit.com/r/shortcuts/comments/9jszwn/import_weight_data_into_health_app_from_csv/ (You'll have to give the shortcut permission to write to Apple Health)
