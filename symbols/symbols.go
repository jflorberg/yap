package symbols

type Icon string

const (
	//✅ Common Log Levels
	Info    Icon = " ℹ️ "
	Warn    Icon = " ⚠️ "
	Error   Icon = " ❌ "
	Debug   Icon = " 🐛 "
	Success Icon = " ✅ "

	//🚀 App Lifecycle / Events
	Shutdown   Icon = " 🛑 "
	Update     Icon = " 🔄 "
	Deploy     Icon = " 📦 "
	Connection Icon = " 🔌 "

	//🔧 Tools / Actions
	Cog    Icon = " ⚙️ "
	Build  Icon = " 🛠️ "
	Fix    Icon = " 🔧 "
	Search Icon = " 🔍 "
	Save   Icon = " 💾 "

	//📊 Status / Monitoring
	Progress Icon = " ⏳ "
	OK       Icon = " 🟢 "
	Idle     Icon = " 🟡 "
	Fail     Icon = " 🔴 "
	Metrics  Icon = " 📊 "
	Logs     Icon = " 📜 "

	//🕵️ Extra Fun
	Rocket Icon = " 🚀 "
	Trace  Icon = " 🕵️ "
	Secret Icon = " 🤫 "
	Fire   Icon = " 🔥 "
	Party  Icon = " 🎉 "
	Magic  Icon = " ✨ "
	Coffee Icon = " ☕ "
)
