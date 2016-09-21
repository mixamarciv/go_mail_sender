package main

import (

	//"math/rand"
	//"bytes"

	"strings"

	"github.com/go-gomail/gomail"
	flags "github.com/jessevdk/go-flags"
	//mf "github.com/mixamarciv/gofncstd3000"
	"strconv"
)

var opts struct {
	// Slice of bool will append 'true' each time the option
	// is encountered (can be set multiple times, like -vvv)
	//Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`

	Host string `long:"host" description:"host" default:"smtp.gmail.com"`
	Port int    `long:"port" description:"port number" default:"25"`
	From string `long:"from" description:"send from mail" default:"mailsendservice1@gmail.com"`
	Pass string `long:"pass" description:"password for sender" default:"AsPeefW2m42i03yqVB9f"`

	To string `long:"to" description:"send email to mails list" default:"mailsendservice1@gmail.com"`

	Subject  string   `long:"subject" description:"Subject" default:""`
	Body     string   `long:"body" description:"text in body mail" default:""`
	Bodytype string   `long:"bodytype" description:"type text in body mail" default:"text/html"`
	Files    []string `long:"files" description:"files for attach"`
}

func main() {
	//var parser = new(flags.Parser)
	parser := flags.NewParser(&opts, 0)
	_, err := parser.Parse()
	LogPrintErrAndExit("ОШИБКА разбора параметров запуска скрипта: ", err)

	{ //отправка почты
		msg := "sending mail: \n"
		msg = msg + "from: " + opts.From + " (" + opts.Host + ":" + strconv.Itoa(opts.Port) + ");\nto: " + opts.To + ";\n"
		msg = msg + "subject: \"" + opts.Subject + "\";\nbodytype: \"" + opts.Bodytype + "\";\n" + "body: \"" + opts.Body + "\";\n"

		m := gomail.NewMessage()
		m.SetHeader("From", opts.From)
		m.SetHeader("To", opts.To)
		m.SetAddressHeader("Cc", "mailsendservice1@gmail.com", "наш мейл")
		m.SetHeader("Subject", opts.Subject)
		m.SetBody(opts.Bodytype, opts.Body)
		//m.Attach("d:\\download\\!images\\vezde_nameki\\vezde_nameki.png")
		//m.Attach("d:\\download\\!images\\vezde_nameki\\vezde_nameki.jpg")
		for i := 0; i < len(opts.Files); i++ {
			msg = msg + "file" + strconv.Itoa(i+1) + ": \"" + opts.Files[i] + "\";\n"
			m.Attach(opts.Files[i])
		}
		LogPrint(msg)

		t := strings.SplitN(opts.From, "@", 2)
		if len(t) != 2 {
			LogPrintAndExit("ОШИБКА не верно задан mail отправителя \"" + opts.From + "\"")
		}
		login := t[0]

		d := gomail.NewPlainDialer(
			//"smtp.gmail.com", 25, "mailsendservice1", "AsPeefW2m42i03yqVB9f")
			opts.Host, opts.Port, login, opts.Pass)

		err := d.DialAndSend(m)

		LogPrintErrAndExit("ОШИБКА отправки сообщения: ", err)
		LogPrint("письмо успешно отправлено")
	}
}
