/*
 * Copyright 2018 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lib

import (
	"github.com/satori/go.uuid"
)

type Response struct {
	Message string `json:"message,omitempty"`
}

type PipelineResponse struct {
	Id uuid.UUID `json:"id,omitempty"`
}

type Pipeline struct {
	Id        string `bson:"id" json:"id"`
	UserId    string
	Operators []Operator `json:"operators,omitempty"`
}

type Operator struct {
	Id          int    `json:"Id,omitempty"`
	Name        string `json:"Name,omitempty"`
	ImageId     string `json:"ImageId,omitempty"`
	InputTopics []InputTopic
}

type InputTopic struct {
	Name        string
	FilterType  string
	FilterValue string
	Mappings    []Mapping
}

type Mapping struct {
	Dest   string
	Source string
}
