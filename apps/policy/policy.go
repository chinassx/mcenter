package policy

import (
	"context"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/pb/resource"

	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcenter/common/format"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// New 新实例
func New(req *CreatePolicyRequest) (*Policy, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	p := &Policy{
		Meta: resource.NewMeta(),
		Spec: req,
	}
	p.genID()

	return p, nil
}

// NewDefaultPolicy todo
func NewDefaultPolicy() *Policy {
	return &Policy{}
}

func (p *Policy) genID() {
	h := fnv.New32a()
	hashedStr := fmt.Sprintf("%s-%s-%s-%s",
		p.Spec.Domain, p.Spec.Namespace, p.Spec.Username, p.Spec.RoleId)

	h.Write([]byte(hashedStr))
	p.Meta.Id = fmt.Sprintf("%x", h.Sum32())
}

// IsAllNamespace 是否是对账所有namespace的测试
func (p *Policy) IsAllNamespace() bool {
	return p.Spec.Namespace == "*"
}

func (p *Policy) ToJson() string {
	return format.Prettify(p)
}

func (u *Policy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*resource.Meta
		*CreatePolicyRequest
		Role *role.Role `json:"role"`
	}{u.Meta, u.Spec, u.Role})
}

// NewCreatePolicyRequest 请求实例
func NewCreatePolicyRequest() *CreatePolicyRequest {
	return &CreatePolicyRequest{
		Enabled: true,
		Scope:   map[string]string{},
		Extra:   map[string]string{},
	}
}

// Validate 校验请求合法
func (req *CreatePolicyRequest) Validate() error {
	return validate.Struct(req)
}

func (p *CreatePolicyRequest) ScopeToString() string {
	list := []string{}
	for k, v := range p.Scope {
		list = append(list, fmt.Sprintf("%s=%s", k, v))
	}

	if len(list) == 0 {
		return ""
	}

	return strings.Join(list, ";")
}

// NewPolicySet todo
func NewPolicySet() *PolicySet {
	return &PolicySet{
		Items: []*Policy{},
	}
}

func (p *PolicySet) ToJson() string {
	return format.Prettify(p)
}

// Users 策略包含的所有用户ID, 已去重
func (s *PolicySet) Users() []string {
	users := map[string]struct{}{}
	for i := range s.Items {
		users[s.Items[i].Spec.Username] = struct{}{}
	}

	set := make([]string, 0, len(users))
	for k := range users {
		set = append(set, k)
	}

	return set
}

// Add 添加
func (s *PolicySet) Add(e *Policy) {
	s.Items = append(s.Items, e)
}

// Length todo
func (s *PolicySet) Length() int {
	return len(s.Items)
}

// GetRoles todo
func (s *PolicySet) GetRoles(ctx context.Context, r role.Service, withPermission bool) (*role.RoleSet, error) {
	set := role.NewRoleSet()
	for i := range s.Items {
		req := role.NewDescribeRoleRequestWithID(s.Items[i].Spec.RoleId)
		ins, err := r.DescribeRole(ctx, req)
		if err != nil {
			return nil, err
		}
		// 继承policy上的范围限制
		ins.Scope = s.Items[i].Spec.ScopeToString()
		set.Add(ins)
	}
	return set, nil
}

// UserRoles 获取用户的角色
func (s *PolicySet) UserRoles(username string) []string {
	rns := []string{}
	for i := range s.Items {
		item := s.Items[i]
		if item.Spec.Username == username {
			rns = append(rns, item.Spec.RoleId)
		}
	}

	if len(rns) == 0 {
		rns = append(rns, "vistor")
	}

	return rns
}

func (s *PolicySet) GetNamespace() (nss []string) {
	nmap := map[string]struct{}{}
	for i := range s.Items {
		nmap[s.Items[i].Spec.Namespace] = struct{}{}
	}

	for k := range nmap {
		nss = append(nss, k)
	}

	return
}

func (s *PolicySet) GetNamespaceWithPage(page *request.PageRequest) (nss []string, total int64) {
	nmap := map[string]struct{}{}
	for i := range s.Items {
		// 如果policy的namespace为* , 表示所有namespace
		if s.Items[i].Spec.Namespace == "*" {
			return []string{}, 0
		}

		nmap[s.Items[i].Spec.Namespace] = struct{}{}
	}

	offset := page.PageSize*page.PageNumber - 1
	end := offset + page.PageSize

	var count uint64 = 0
	for k := range nmap {
		if count >= offset && count < end {
			nss = append(nss, k)
		}

		count++
	}

	return nss, int64(len(nmap))
}

// Validate todo
func (req *DescribePolicyRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("policy id required")
	}

	return nil
}

// Validate 校验请求是否合法
func (req *QueryPolicyRequest) Validate() error {
	return validate.Struct(req)
}

// NewDescriptPolicyRequest new实例
func NewDescriptPolicyRequest() *DescribePolicyRequest {
	return &DescribePolicyRequest{}
}

// NewDeletePolicyRequestWithID todo
func NewDeletePolicyRequestWithID(id string) *DeletePolicyRequest {
	return &DeletePolicyRequest{Id: id}
}

// NewQueryPolicyRequest 列表查询请求
func NewQueryPolicyRequest() *QueryPolicyRequest {
	return &QueryPolicyRequest{
		Page:          request.NewDefaultPageRequest(),
		WithRole:      false,
		WithNamespace: false,
	}
}

// NewDeletePolicyRequestWithNamespaceID todo
func NewDeletePolicyRequestWithNamespace(domain, namespace string) *DeletePolicyRequest {
	return &DeletePolicyRequest{
		Domain:    domain,
		Namespace: namespace,
	}
}
