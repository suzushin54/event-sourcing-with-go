package pkg

import "github.com/bwmarrin/snowflake"

type SnowflakeNode interface {
	Generate() snowflake.ID
}

func NewSnowflakeNode() (SnowflakeNode, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	return node, nil
}
