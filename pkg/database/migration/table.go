/**
 * Created by zc on 2020/6/3.
 */
package migration

import "stone/pkg/database/migration/data"

func tables() []interface{} {
	return []interface{}{
		&data.User{},
		&data.Space{}, &data.SpaceAccess{}, &data.SpaceRule{},
		&data.Config{}, &data.ConfigRecord{}, &data.ConfigRuleRelate{},
		&data.App{},
		&data.Rule{}, &data.RuleRecord{},
	}
}