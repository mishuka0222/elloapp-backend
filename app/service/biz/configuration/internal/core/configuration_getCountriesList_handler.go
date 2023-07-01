package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/configuration/configuration"
)

// ConfigurationGetCountryList
// help.getCountriesList#735787a8 lang_code:string hash:int = help.CountriesList;
func (c *ConfigurationCore) ConfigurationGetCountryList(in *configuration.GetCountryListReq) (*configuration.GetCountryListResp, error) {

	countriesData, err := c.svcCtx.Dao.CountriesDAO.SelectList(c.ctx)
	if err != nil {
		return nil, err
	}

	var countries []*configuration.CountryDO

	for _, item := range countriesData {
		con := &configuration.CountryDO{
			Name: item.Name,
			Flag: item.Flag,
			Code: item.Code,
		}
		countries = append(countries, con)
	}

	return &configuration.GetCountryListResp{Data: countries}, nil
}
