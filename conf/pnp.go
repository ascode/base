/**
阿里隐私号码保护配置
 */
package conf

type PnP struct {
	RegionId     		string `toml:"region_id"`
	AccessKeyId      	string `toml:"access_key_id"`
	AccessKeySecret     string `toml:"access_key_secret"`
	Domain				string `toml:"domain"`
	Version				string `toml:"version"`
	PoolKey				string `toml:"pool_key"`
}