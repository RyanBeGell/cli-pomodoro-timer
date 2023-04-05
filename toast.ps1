param(
    [string]$title,
    [string]$message,
    [string]$imagePath
)

[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null;
[Windows.UI.Notifications.ToastNotification, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null;
[Windows.Data.Xml.Dom.XmlDocument, Windows.Data.Xml.Dom.XmlDocument, ContentType = WindowsRuntime] | Out-Null;

$template = @"
<toast>
<visual>
<binding template="ToastGeneric">
<image placement="appLogoOverride" src="$imagePath"/>
<text><![CDATA[$title]]></text>
<text><![CDATA[$message]]></text>
</binding>
</visual>
</toast>
"@

$xml = New-Object Windows.Data.Xml.Dom.XmlDocument
$xml.LoadXml($template)

$toast = New-Object Windows.UI.Notifications.ToastNotification $xml
$toast.Tag = "PowerShell"
$toast.Group = "PowerShell"
$toast.ExpirationTime = [DateTimeOffset]::Now.AddSeconds(5)

$appId = "Microsoft.WindowsTerminal_8wekyb3d8bbwe!App"
$notifier = [Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier($appId)
$notifier.Show($toast)
