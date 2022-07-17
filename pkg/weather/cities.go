package weather

// City must be stored privately and set by telegram location message

type City interface {
	Lat() string
	Lon() string
	NameIn() string
}

type Moscow struct{}

func (m *Moscow) Lat() string {
	return "55.734180"
}
func (m *Moscow) Lon() string {
	return "37.588218"
}
func (m *Moscow) NameIn() string {
	return "Москве"
}

type Kazan struct{}

func (k *Kazan) Lat() string {
	return "55.783956"
}
func (k *Kazan) Lon() string {
	return "49.127972"
}
func (k *Kazan) NameIn() string {
	return "Казани"
}

type Ryazan struct{}

func (r *Ryazan) Lat() string {
	return "54.730411"
}
func (r *Ryazan) Lon() string {
	return "39.840408"
}
func (r *Ryazan) NameIn() string {
	return "Рязани"
}
