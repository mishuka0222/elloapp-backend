package mtproto

// MakeGeoPointByInput
// inputGeoPointEmpty#e4c123d6 = InputGeoPoint;
// inputGeoPoint#f3b7acc9 lat:double long:double = InputGeoPoint;
//
// geoPointEmpty#1117dd5f = GeoPoint;
// geoPoint#296f104 long:double lat:double access_hash:long = GeoPoint;
func MakeGeoPointByInput(geoPoint *InputGeoPoint) (geo *GeoPoint) {
	switch geoPoint.PredicateName {
	case Predicate_inputGeoPointEmpty:
		geo = MakeTLGeoPointEmpty(nil).To_GeoPoint()
	case Predicate_inputGeoPoint:
		// TODO(@benqi): check access_hash
		geo = MakeTLGeoPoint(&GeoPoint{
			Long: geoPoint.Long,
			Lat:  geoPoint.Lat,
		}).To_GeoPoint()
	default:
		// log.Errorf("invalid predicateName - %v", geoPoint)
	}
	return
}
