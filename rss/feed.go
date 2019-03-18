package rss

import "log"

var Feeds = []struct {
	Name     string
	Category string
	Link     string
	Kind     string
	Enabled   bool
	Load     bool
}{ // A. Material Vivo Animal y Vegetal
	{
		"Material vivo vegetal y animal, accesorios y suministros",
		"A. Material Vivo Animal y Vegetal",
		"https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-10000000.xml",
		"rss",
		true,
		true,
	},
	// B.  Materias Primas
	{
		"Material mineral, textil y vegetal y animal no comestible",
		"B. Materias Primas",
		"https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-11000000.xml",
		"rss",
		true,
		true,
	},
	{
		"Material químico incluyendo bioquímicos y materiales de gas",
		"B. Materias Primas",
		"https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-12000000.xml",
		"rss",
		true,
		true,
	},
	{
		"Materiales de resina, colofonia, caucho, espuma, película y elastómericos",
		"B. Materias Primas",
		"https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-13000000.xml",
		"rss",
		true,
		true,
	},
	{
		"Materiales y productos de papel",
		"B. Materias Primas",
		"https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-14000000.xml",
		"rss",
		true,
		true,
	},
	{
		"Materiales combustibles, aditivos para combustibles, lubricantes y anticorrosivos",
		"B. Materias Primas",
		"https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-15000000.xml",
		"rss",
		true,
		true,
	},
	// C. Maquinaria, Herramientas, Equipo Industrial y Vehículos
	{
		"Maquinaria y accesorios de minería y perforación de pozos",
		"C. Maquinaria, Herramientas, Equipo Industrial y Vehículos",
		"https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-20000000.xml",
		"rss",
		true,
		true,
	},
	{
		"Maquinaria y accesorios para agricultura, pesca, silvicultura y fauna",
		"C. Maquinaria, Herramientas, Equipo Industrial y Vehículos",
		"https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-21000000.xml",
		"rss",
		true,
		true,
	},
	{
		"Maquinaria y accesorios para construcción y edificación",
		"C. Maquinaria, Herramientas, Equipo Industrial y Vehículos",
		"https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-22000000.xml",
		"rss",
		true,
		true,
	},
	{
		"Maquinaria y accesorios para manufactura y procesamiento industrial",
		"C. Maquinaria, Herramientas, Equipo Industrial y Vehículos",
		"https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-23000000.xml",
		"rss",
		true,
		true,
	},
	{
		"Maquinaria, accesorios y suministros para manejo, acondicionamiento y almacenamiento de materiales",
		"C. Maquinaria, Herramientas, Equipo Industrial y Vehículos",
		"https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-24000000.xml",
		"rss",
		true,
		true,
	},
	{
		"Vehículos comerciales, militares y particulares, accesorios y componentes",
		"C. Maquinaria, Herramientas, Equipo Industrial y Vehículos",
		"https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-25000000.xml",
		"rss",
		true,
		true,
	},
	{
		"Maquinaria y accesorios para generación y distribución de energía",
		"C. Maquinaria, Herramientas, Equipo Industrial y Vehículos",
		"https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-26000000.xml",
		"rss",
		true,
		true,
	},
	{
		"Herramientas y maquinaria general",
		"C. Maquinaria, Herramientas, Equipo Industrial y Vehículos",
		"https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-27000000.xml",
		"rss",
		true,
		true,
	},
}

func ReadFeeds() {
	for _, source := range Feeds {
		if source.Enabled {
			err:= getDocument(source.Link, true)
			if err!=nil{
				log.Println("Azuquita pa' el café")
			}
		}

	}
}
