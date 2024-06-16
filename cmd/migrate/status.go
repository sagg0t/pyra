package main

// const (
// 	separator = "-----------------+-------+--------------------------------+---------------------+--------"
// 	rowFormat = "%s | %5s | %30s | %s | %s\n"
// )
//
// var statusHeader string = fmt.Sprintf("%16s | %s | %30s | %19s | %7s\n%s", "Version", "State", "Name", "Applied At", "Files", separator)
//
// func status() {
// 	migs, err := engine.Status()
// 	if err != nil {
// 		slog.Error("error whil retrieving migration status", "error", err)
// 		os.Exit(1)
// 	}
//
// 	sort.Sort(sort.Reverse(migs))
//
// 	fmt.Println(statusHeader)
//
// 	for _, m := range migs {
// 		state := "up"
// 		if m.AppliedAt.Equal(time.Time{}) {
// 			state = "down"
// 		}
//
// 		var filesPresence string
// 		if m.UpFile != "" && m.DownFile != "" {
// 			filesPresence = "up/down"
// 		} else if m.UpFile != "" {
// 			filesPresence = "up"
// 		} else if m.DownFile != "" {
// 			filesPresence = "down"
// 		} else {
// 			filesPresence = "NONE"
// 		}
//
// 		fmt.Printf(rowFormat, m.Version, state, m.Name, m.AppliedAt.Format(time.DateTime), filesPresence)
// 	}
// }
