package util

import "net/url"

func URLParseQueryPatch(u string, p url.Values) (*url.URL, error) {
	purl, err := url.Parse(u)
	if err != nil {
		return purl, err
	}
	return URLQueryPatch(purl, p)
}
func URLQueryPatch(purl *url.URL, p url.Values) (*url.URL, error) {
	query := purl.Query() // just a copy
	for k, vs := range p {
		for _, v := range vs {
			query.Add(k, v)
		}
	}
	purl.RawQuery = query.Encode() // re-assign
	return purl, nil
}
