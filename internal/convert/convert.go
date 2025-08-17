package convert

// ToServices converts a config.File services to a slice of runner.Service.
//func ToServices(f *config.File) ([]runner.Service, error) {
//	out := make([]runner.Service, 0, len(f.Services))
//	for name, s := range f.Services {
//		rs := runner.Service{
//			Name:      name,
//			Image:     s.Image,
//			Env:       map[string]string{},
//			DependsOn: append([]string{}, s.DependsOn...),
//		}
//		for k, v := range s.Env {
//			rs.Env[k] = v
//		}
//
//		for _, p := range s.Ports {
//			h, c, ok := strings.Cut(p, ":")
//			if !ok {
//				return nil, fmt.Errorf("service %s: invalid port %q", name, p)
//			}
//			hh, _ := strconv.Atoi(h)
//			cc, _ := strconv.Atoi(c)
//			rs.Ports = append(rs.Ports, runner.Port{Host: hh, Container: cc})
//		}
//		if len(s.Command) > 0 {
//			rs.Command = append([]string{}, s.Command...)
//		}
//		if s.Build != nil {
//			rs.Build = &runner.BuildSpec{
//				Context:    s.Build.Context,
//				Dockerfile: s.Build.Dockerfile,
//				Args:       s.Build.Args,
//			}
//		}
//		if s.Healthcheck != nil {
//			rs.Healthcheck = &runner.HealthSpec{
//				HTTPGet:   s.Healthcheck.HTTPGet,
//				TCPSocket: s.Healthcheck.TCPSocket,
//				Cmd:       s.Healthcheck.Cmd,
//			}
//		}
//		out = append(out, rs)
//	}
//	return out, nil
//}
//
