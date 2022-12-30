package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/configuration/configuration"
)

// HelpGetCountriesList
// help.getCountriesList#735787a8 lang_code:string hash:int = help.CountriesList;
func (c *ConfigurationCore) HelpGetCountriesList(in *mtproto.TLHelpGetCountriesList) (*mtproto.Help_CountriesList, error) {

	countriesData, err := c.svcCtx.Dao.ConfigurationClient.ConfigurationGetCountryList(c.ctx, &configuration.GetCountryListReq{LangCode: in.LangCode})
	if err != nil {
		return nil, err
	}

	var countries []*mtproto.Help_Country

	for _, item := range countriesData.Data {
		con := (&mtproto.Help_Country{
			DefaultName: item.Name,
			Iso2:        item.Flag,
			CountryCodes: []*mtproto.Help_CountryCode{
				(&mtproto.Help_CountryCode{CountryCode: item.Code}).
					To_HelpCountryCode().To_Help_CountryCode(),
			},
		}).To_HelpCountry().To_Help_Country()
		countries = append(countries, con)
	}

	return mtproto.MakeTLHelpCountriesList(&mtproto.Help_CountriesList{
		Countries: countries,
		Hash:      in.Hash,
	}).To_Help_CountriesList(), nil
}
