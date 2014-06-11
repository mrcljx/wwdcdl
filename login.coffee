casper = require("casper").create()
fs = require "fs"
system = require "system"

url = "https://developer.apple.com/membercenter/"

casper.start url, ->
  casper.log "[URL] #{@getCurrentUrl()}", "debug"
  casper.log "Filling out form...", "debug"

  theAccountName = String(casper.cli.args[1] ? "")

  unless theAccountName.length > 0
    system.stdout.write "Username: "
    theAccountName = system.stdin.readLine()

  theAccountPW = String(casper.cli.args[2] ? "")

  unless theAccountPW.length > 0
    system.stdout.write "Password: "
    theAccountPW = system.stdin.readLine()

  @fill "form[name=appleConnectForm]", { theAccountName, theAccountPW }, true

casper.waitFor ->
  @getCurrentUrl().indexOf("daw.apple.com") < 0
, ->
  casper.log "[URL] #{@getCurrentUrl()}", "debug"
  
  unless @exists "form#saveTeamSelection"
    casper.log "No team selection necessary.", "debug"
    return
    
  options = @evaluate ->
    for option in document.querySelectorAll("form#saveTeamSelection option")
    	{ id: option.value, text: option.innerText }

  for option in options
    if casper.cli.args[3] is option.id
      selectedOption = option
      break

  while !selectedOption?
    @echo ""

    for option, index in options
      @echo "#{index+1}) #{option.text} (#{option.id})"

    @echo ""
    @echo "Please choose a team [1-#{options.length}]: "

    index = system.stdin.readLine()
    selectedOption = options[Number(index) - 1]
  
  casper.log "Selecting team: #{selectedOption.text}", "info"

  @fill "form#saveTeamSelection", { memberDisplayId: selectedOption.id }, true
, ->
  @die "Invalid credentials."

casper.waitFor ->
  @getCurrentUrl().indexOf("selectTeam.action") < 0
, ->
  casper.log "[URL] #{@getCurrentUrl()}", "debug"
, ->
  @die "Invalid team."

casper.wait 250, ->
  cookies = JSON.stringify casper.page.cookies
  file = casper.cli.args[0] ? "cookies.json"  
  fs.write file, cookies, "wb"
  casper.log "Cookies written to file: #{file}", "debug"
  casper.log "Exiting...", "debug"

casper.run()