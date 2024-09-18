/**
 * Copyright 2014 @ 56x.net.
 * name :
 * author : jarryliu
 * date : 2013-12-03 23:20
 * description :
 * history :
 */
package query

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ixre/go2o/core/domain/interface/member"
	"github.com/ixre/go2o/core/domain/interface/wallet"
	"github.com/ixre/go2o/core/dto"
	"github.com/ixre/go2o/core/infrastructure/fw"
	"github.com/ixre/go2o/core/service/proto"
	"github.com/ixre/gof/db"
	"github.com/ixre/gof/db/orm"
)

type MemberQuery struct {
	db.Connector
	o orm.Orm
	fw.BaseRepository[member.Member]
	balanceLog  fw.BaseRepository[member.BalanceLog]
	integralLog fw.BaseRepository[member.IntegralLog]
	walletLog   fw.BaseRepository[wallet.WalletLog]
	levelRepo   fw.BaseRepository[member.Level]
	certRepo    fw.BaseRepository[member.CerticationInfo]
}

func NewMemberQuery(o orm.Orm, fo fw.ORM) *MemberQuery {
	q := &MemberQuery{
		Connector: o.Connector(),
		o:         o}
	q.ORM = fo
	q.balanceLog.ORM = fo
	q.integralLog.ORM = fo
	q.walletLog.ORM = fo
	q.levelRepo.ORM = fo
	q.certRepo.ORM = fo
	return q
}

// 获取会员列表
func (m *MemberQuery) QueryMemberList(ids []int64) []*dto.MemberSummary {
	var list []*dto.MemberSummary
	strIds := make([]string, len(ids))
	for i, v := range ids {
		strIds[i] = strconv.Itoa(int(v))
	}
	if len(ids) > 0 {
		inStr := strings.Join(strIds, ",") // order by field(field,val1,val2,val3)按IN的顺序排列
		query := fmt.Sprintf(`SELECT m.id,m.user,m.nick_name,m.profile_photo,m.exp,m.level,
				lv.name as level_name,a.integral,a.balance,a.wallet_balance,
				a.grow_balance,a.grow_amount,a.grow_earnings,a.grow_total_earnings,
				m.update_time FROM mm_member m INNER JOIN mm_level lv
				ON m.level = lv.id INNER JOIN mm_account a ON
				 a.member_id = m.id AND m.id IN(%s) order by field(m.id,%s)`, inStr, inStr)
		m.o.SelectByQuery(&list, query)
	}
	return list
}

// 获取账户余额分页记录
func (m *MemberQuery) PagedBalanceAccountLog(memberId int64, valueFilter int32, begin, end int,
	where, orderBy string) (num int, rows []*proto.SMemberAccountLog) {
	d := m.Connector
	if valueFilter == 1 {
		where += ` AND change_value > 0 `
	}
	if valueFilter == 2 {
		where += ` AND change_value < 0 `
	}
	if orderBy != "" {
		orderBy = "ORDER BY " + orderBy
	}
	d.ExecScalar(fmt.Sprintf(`SELECT COUNT(1) FROM mm_balance_log
			WHERE member_id= $1 %s`, where), &num, memberId)
	sqlLine := fmt.Sprintf(`SELECT id,kind,subject,outer_no,change_value,procedure_fee,
	balance,review_status,create_time FROM mm_balance_log
			WHERE member_id= $1 %s %s LIMIT $3 OFFSET $2`,
		where, orderBy)
	err := d.Query(sqlLine, func(_rows *sql.Rows) {
		for _rows.Next() {
			e := proto.SMemberAccountLog{}
			_rows.Scan(&e.Id, &e.Kind, &e.Subject, &e.OuterNo,
				&e.Value, &e.TransactionFee, &e.Balance,
				&e.ReviewStatus, &e.CreateTime)
			rows = append(rows, &e)
		}
	}, memberId, begin, end-begin)
	if err != nil {
		log.Println("----", err)
	}

	return num, rows
}

// 获取账户余额分页记录
func (m *MemberQuery) PagedIntegralAccountLog(memberId int64, valueFilter int32,
	begin, over int64, sortBy string) (num int, rows []*proto.SMemberAccountLog) {
	d := m.Connector
	where := ""
	if valueFilter == 1 {
		where += ` AND change_value > 0 `
	}
	if valueFilter == 2 {
		where += ` AND change_value < 0 `
	}
	d.ExecScalar(fmt.Sprintf(`SELECT COUNT(1) FROM mm_integral_log 
			WHERE member_id= $1 %s`, where), &num, memberId)
	if num > 0 {
		orderBy := ""
		if sortBy != "" {
			orderBy = "ORDER BY " + sortBy
		}
		sqlLine := fmt.Sprintf(`SELECT id,kind,subject,outer_no,change_value,
		balance,review_status,create_time FROM mm_integral_log 
			WHERE member_id= $1 %s %s LIMIT $3 OFFSET $2`, where, orderBy)
		err := d.Query(sqlLine, func(_rows *sql.Rows) {
			for _rows.Next() {
				e := proto.SMemberAccountLog{}
				_rows.Scan(&e.Id, &e.Kind, &e.Subject, &e.OuterNo,
					&e.Value, &e.Balance,
					&e.ReviewStatus, &e.CreateTime)
				rows = append(rows, &e)
			}
		}, memberId, begin, over-begin)
		if err != nil {
			log.Println("[ GO2O][ Params]: query error ", err.Error())
		}
	}
	return num, rows
}

// 获取账户余额分页记录
func (m *MemberQuery) PagedWalletAccountLog(memberId int64, valueFilter int32, begin, end int,
	where, orderBy string) (num int, rows []*proto.SMemberAccountLog) {
	d := m.Connector
	if valueFilter == 1 {
		where += ` AND change_value > 0 `
	}
	if valueFilter == 2 {
		where += ` AND change_value < 0 `
	}
	if orderBy != "" {
		orderBy = "ORDER BY " + orderBy
	}
	var walletId = 0
	d.ExecScalar(`SELECT id FROM wal_wallet LEFT JOIN mm_account 
	ON mm_account.wallet_code=wal_wallet.hash_code
	WHERE mm_account.member_id=$1 limit 1`, &walletId, memberId)

	d.ExecScalar(fmt.Sprintf(`SELECT COUNT(1) FROM wal_wallet_log WHERE wallet_id = $1 %s`, where), &num, walletId)

	//rows = make([]*proto.SMemberAccountLog,0)

	if num > 0 {
		cmd := fmt.Sprintf(`SELECT id,kind,subject,outer_no,change_value,procedure_fee,
			balance,review_status,create_time FROM wal_wallet_log 
			WHERE wallet_id = $1 %s %s LIMIT $3 OFFSET $2`,
			where, orderBy)
		err := d.Query(cmd, func(_rows *sql.Rows) {
			for _rows.Next() {
				e := proto.SMemberAccountLog{}
				_rows.Scan(&e.Id, &e.Kind, &e.Subject, &e.OuterNo,
					&e.Value, &e.TransactionFee, &e.Balance,
					&e.ReviewStatus, &e.CreateTime)
				rows = append(rows, &e)
			}
		}, walletId, begin, end-begin)
		if err != nil {
			log.Println("[ GO2O][ ERROR]: query wallet log", err.Error(), ";sql=", cmd)
		}
	}
	return num, rows
}

// 获取最近的余额变动信息
func (m *MemberQuery) GetLatestWalletLogByKind(memberId int64, kind int) *member.WalletAccountLog {
	var info = new(member.WalletAccountLog)
	if err := m.o.GetBy(info, "member_id= $1 AND kind= $2 ORDER BY create_time DESC",
		memberId, kind); err == nil {
		return info
	}
	return nil
}

// 筛选会员根据用户或者手机
func (m *MemberQuery) FilterMemberByUserOrPhone(key string) []*dto.SimpleMember {
	qp := "%" + key + "%"
	list := make([]*dto.SimpleMember, 0)
	var id int
	var user, name, phone, portrait string
	m.Query(`SELECT id,user,mm_profile.name,mm_profile.phone,
        mm_profile.profile_photo FROM mm_member
        INNER JOIN mm_profile ON mm_profile.member_id=mm_member.id
        WHERE user LIKE $1 OR mm_profile.name LIKE $2 OR
        mm_profile.phone LIKE $3`, func(rows *sql.Rows) {
		for rows.Next() {
			rows.Scan(&id, &user, &name, &phone, &portrait)
			list = append(list, &dto.SimpleMember{
				Id:     id,
				User:   user,
				Name:   name,
				Phone:  phone,
				Avatar: portrait,
			})
		}
	}, qp, qp, qp)
	return list
}

func (m *MemberQuery) GetMemberByUserOrPhone(key string) *dto.SimpleMember {
	e := dto.SimpleMember{}
	err := m.QueryRow(`SELECT id,user,mm_profile.name,mm_profile.phone,
        mm_profile.profile_photo FROM mm_member
        INNER JOIN mm_profile ON mm_profile.member_id=mm_member.id
        WHERE user = $1 OR mm_profile.phone = $2`, func(rows *sql.Row) error {
		er := rows.Scan(&e.Id, &e.User, &e.Name, &e.Phone, &e.Avatar)
		return er
	}, key, key)
	if err == nil {
		return &e
	}
	return nil
}

// 根据手机获取会员编号
func (m *MemberQuery) GetMemberIdByPhone(phone string) int64 {
	var id int64
	m.ExecScalar(`SELECT id FROM mm_member
        INNER JOIN mm_profile ON mm_profile.member_id=mm_member.id
        WHERE mm_profile.phone = $1 LIMIT 1`, &id, phone)
	return id
}

// 会员推广排名
func (m *MemberQuery) GetMemberInviRank(mchId int64, allTeam bool, levelComp string, level int,
	startTime int64, endTime int64, num int) []*dto.RankMember {
	list := make([]*dto.RankMember, 0)
	var id int64
	var user, name string
	var inviNum, totalNum, regTime int
	var rank = 0

	var sortField = "t.all_num DESC"
	if !allTeam {
		sortField = "t.invi_num DESC"
	}

	var levelCompStr = fmt.Sprintf("%s%d", levelComp, level)
	//{level_comp}{level_value}

	m.Query(fmt.Sprintf(`SELECT id,user,name,invi_num,all_num,reg_time FROM ( SELECT m.*,
 (SELECT COUNT(1) FROM mm_relation r INNER JOIN mm_member m1 ON m1.id = r.member_id WHERE
  (m1.level%s) AND r.inviter_id = m.id
	AND r.reg_mch_id=rl.reg_mch_id  AND m1.reg_time BETWEEN
  $1 AND $2 ) as invi_num,
	((SELECT COUNT(1) FROM mm_relation r INNER JOIN mm_member m1 ON m1.id = r.member_id WHERE
  (m1.level%s) AND r.inviter_id = m.id
	AND r.reg_mch_id=rl.reg_mch_id AND m1.reg_time BETWEEN
  $3 AND $4 )+
 (SELECT COUNT(1) FROM mm_relation r INNER JOIN mm_member m1
  ON m1.id = r.member_id WHERE (m1.level%s) AND inviter_id IN
	(SELECT member_id FROM mm_relation r INNER JOIN mm_member m1 ON m1.id = r.member_id WHERE
  (m1.level%s) AND r.inviter_id =
    m.id AND r.reg_mch_id=rl.reg_mch_id AND m1.reg_time BETWEEN
  $5 AND $6 ))) as all_num
 FROM mm_member m INNER JOIN mm_relation rl ON m.id= rl.member_id
 WHERE rl.reg_mch_id = $7 AND state= $8) t ORDER BY %s,t.reg_time asc
 LIMIT $9`, levelCompStr, levelCompStr, levelCompStr, levelCompStr, sortField), func(rows *sql.Rows) {
		for rows.Next() {
			rows.Scan(&id, &user, &name, &inviNum, &totalNum, &regTime)
			rank++
			list = append(list, &dto.RankMember{
				Id:       id,
				Usr:      user,
				Name:     name,
				RankNum:  rank,
				InviNum:  inviNum,
				TotalNum: totalNum,
				RegTime:  regTime,
			})
		}
	}, startTime, endTime, startTime, endTime, startTime, endTime, mchId, 1, num)

	return list
}

// 获取分页店铺收藏
func (m *MemberQuery) PagedShopFav(memberId int64, begin, end int,
	where string) (num int, rows []*dto.PagedShopFav) {
	d := m.Connector
	if len(where) > 0 {
		where = " AND " + where
	}
	d.ExecScalar(fmt.Sprintf(`SELECT COUNT(1) FROM mm_favorite f
	INNER JOIN  mch_shop s ON f.refer_id =s.id
    INNER JOIN mch_online_shop o ON s.id = o.shop_id
    INNER JOIN mch_merchant mch ON mch.id = s.vendor_id
    WHERE f.member_id= $1 AND f.fav_type= $2 %s`, where), &num,
		memberId, member.FavTypeShop)

	if num > 0 {
		sqlLine := fmt.Sprintf(`SELECT f.id,s.id as shop_id,mch.id as mch_id,
    s.shop_name,o.logo,f.update_time FROM mm_favorite f
    INNER JOIN  mch_shop s ON f.refer_id =s.id
    INNER JOIN mch_online_shop o ON s.id = o.shop_id
    INNER JOIN mch_merchant mch ON mch.id = s.vendor_id
    WHERE f.member_id= $1 AND f.fav_type= $2 %s ORDER BY f.update_time DESC LIMIT $4 OFFSET $3`,
			where)
		d.Query(sqlLine, func(rs *sql.Rows) {
			for rs.Next() {
				e := dto.PagedShopFav{}
				rs.Scan(&e.Id, &e.ShopId, &e.MchId, &e.ShopName,
					&e.Logo, &e.UpdateTime)
				rows = append(rows, &e)
			}
		}, memberId, member.FavTypeShop, begin, end-begin)
	} else {
		rows = make([]*dto.PagedShopFav, 0)
	}
	return num, rows
}

// 获取分页店铺收藏
func (m *MemberQuery) PagedGoodsFav(memberId int64, begin, end int,
	where string) (num int, rows []*dto.PagedGoodsFav) {
	d := m.Connector
	if len(where) > 0 {
		where = " AND " + where
	}
	d.ExecScalar(fmt.Sprintf(`SELECT COUNT(1) FROM mm_favorite f
    INNER JOIN item_info gs ON gs.id = f.refer_id
    INNER JOIN product product ON gs.product_id=product.id
    WHERE f.member_id= $1 AND f.fav_type= $2 %s`, where), &num,
		memberId, member.FavTypeGoods)

	if num > 0 {
		sqlLine := fmt.Sprintf(`SELECT f.id,gs.id as goods_id,product.name as goods_name,
            img,sale_price,gs.stock_num,product.update_time
            FROM mm_favorite f INNER JOIN item_info gs ON gs.id = f.refer_id
            INNER JOIN product product ON gs.product_id=product.id
            WHERE f.member_id= $1 AND f.fav_type= $2 %s ORDER BY f.update_time DESC
            LIMIT $4 OFFSET $3`,
			where)
		d.Query(sqlLine, func(rs *sql.Rows) {
			for rs.Next() {
				e := dto.PagedGoodsFav{}
				rs.Scan(&e.Id, &e.SkuId, &e.GoodsName, &e.Image, &e.SalePrice,
					&e.StockNum, &e.UpdateTime)
				rows = append(rows, &e)
			}
		}, memberId, member.FavTypeGoods, begin, end-begin)

	} else {
		rows = make([]*dto.PagedGoodsFav, 0)
	}
	return num, rows
}

// 获取从指定时间到现在推荐指定等级会员的数量
func (m *MemberQuery) GetInviteQuantity(memberId int64, where string) int32 {
	var total int32
	m.Connector.ExecScalar(`SELECT COUNT(1) FROM mm_relation
        INNER JOIN mm_member ON mm_member.id = mm_relation.member_id
        LEFT JOIN mm_cert_info mt ON mt.member_id=mm_member.id
        WHERE inviter_id = $1 `+where, &total, memberId)
	return total
}

// 获取从指定时间到现在推荐指定等级会员的数量
func (m *MemberQuery) GetInviteArray(memberId int64, where string) []int64 {
	arr := []int64{}
	m.Connector.Query(`SELECT mm_relation.member_id FROM mm_relation
        INNER JOIN mm_member ON mm_member.id = mm_relation.member_id
        LEFT JOIN mm_cert_info mt ON mt.member_id=mm_member.id
        WHERE inviter_id = $1 `+where, func(rows *sql.Rows) {
		var i int64
		for rows.Next() {
			if err := rows.Scan(&i); err == nil && i > 0 {
				arr = append(arr, i)
			}
		}
	}, memberId)
	return arr
}

// 获取邀请会员数量
func (m *MemberQuery) InviteMembersQuantity(memberId int64, where string) int {
	total := 0
	err := m.Connector.ExecScalar(`SELECT COUNT(1) FROM mm_relation WHERE `+where, &total)
	if err != nil {
		log.Printf("[ Params][ Error]: invite member quantity error:%s", err.Error())
	}
	return total
}

// 获取订单状态及其数量
func (m *MemberQuery) OrdersQuantity(memberId int64) map[int]int {
	mp := make(map[int]int, 0)
	err := m.Connector.Query(`
	SELECT o.status,COUNT(0) as n FROM sale_sub_order o
	GROUP BY o.status,o.buyer_id,is_forbidden,break_status
	having is_forbidden = 0 AND break_status <> 0 
	 AND o.buyer_id = $1`, func(rows *sql.Rows) {
		s, n := 0, 0
		for rows.Next() {
			rows.Scan(&s, &n)
			mp[s] = n
		}
	}, memberId)
	if err != nil {
		log.Println("[ GO2O][ ERROR]: query order quantity failed! ", err.Error())
	}
	return mp
}

// QueryPagingMemberList 查询会员列表
func (m *MemberQuery) QueryPagingMemberList(p *fw.PagingParams) (*fw.PagingResult, error) {
	tables := `mm_member m LEFT JOIN mm_relation r ON m.id = r.member_id
LEFT JOIN mm_account ac ON m.id = ac.member_id
LEFT JOIN mm_profile pro ON pro.member_id = m.id`

	fields := `
	distinct(m.id),m.nickname,m.real_name,m.username,m.exp,m.profile_photo,pro.gender,pro.birthday,
	m.phone,m.level,m.user_flag,ac.integral,ac.balance,
	ac.total_pay,ac.wallet_balance,m.reg_from,m.reg_time,m.login_time,
	(SELECT id FROM mm_member m2 WHERE m2.id = r.inviter_id) as inviter_id
`
	return fw.UnifinedQueryPaging(m.ORM, p, tables, fields)
}

// QueryPagingStaffs 查询商户员工列表
func (m *MemberQuery) QueryPagingStaffs(p *fw.PagingParams) (*fw.PagingResult, error) {
	tables := `mm_member m
		INNER JOIN mch_staff s ON s.member_id=m.id
		INNER JOIN mch_merchant mch ON mch.id = s.mch_id
		LEFT JOIN mm_profile pro ON pro.member_id = m.id`
	fields := `
	distinct(m.id),m.nickname,m.real_name,m.username,m.exp,m.profile_photo,pro.gender,pro.birthday,
	m.phone,m.level,m.user_flag,
	m.reg_from,m.reg_time,m.login_time,
	s.certified_name,s.is_certified,mch.mch_name
	`
	return fw.UnifinedQueryPaging(m.ORM, p, tables, fields)
}

// QueryMemberBlockList 查询会员拉黑列表
func (m *MemberQuery) QueryMemberBlockList(p *fw.PagingParams) (*fw.PagingResult, error) {
	tables := `mm_block_list b INNER JOIN mm_member m ON b.block_member_id = m.id`
	fields := `b.id,
	b.block_member_id,
	b.block_flag,
	m.profile_photo,
	m.nickname,
	m.username,
	b.create_time`
	return fw.UnifinedQueryPaging(m.ORM, p, tables, fields)
}

func (m *MemberQuery) QueryPagingMemberBalanceLogs(memberId int, p *fw.PagingParams) (*fw.PagingResult, error) {
	p.Equal("member_id", memberId)
	return m.balanceLog.QueryPaging(p)
}

func (m *MemberQuery) QueryPagingMemberIntegralLogs(memberId int, p *fw.PagingParams) (*fw.PagingResult, error) {
	p.Equal("member_id", memberId)
	return m.integralLog.QueryPaging(p)
}

func (m *MemberQuery) QueryPagingMemberWalletLogs(memberId int, p *fw.PagingParams) (*fw.PagingResult, error) {
	var walletId int
	m.ORM.Raw("SELECT id FROM wal_wallet WHERE user_id=? AND wallet_type=? ", memberId, wallet.TPerson).Scan(&walletId)
	if walletId == 0 {
		return nil, fmt.Errorf("user no such wallet,id=%d", memberId)
	}
	p.Equal("wallet_id", walletId)
	return m.walletLog.QueryPaging(p)
}

// QueryPagingCertifications 查询会员认证列表
func (m *MemberQuery) QueryPagingCertifications(p *fw.PagingParams) (*fw.PagingResult, error) {
	tables := `mm_cert_info ti INNER JOIN mm_member m ON m.id=ti.member_id`
	fields := `m.id,ti.member_id,ti.card_type, m.username,m.nickname,m.phone,ti.real_name,
ti.review_status,ti.remark,ti.update_time,ti.card_id`
	return fw.UnifinedQueryPaging(m.ORM, p, tables, fields)

}

// QueryPagingLevels 查询会员等级列表
func (m *MemberQuery) QueryPagingLevels(p *fw.PagingParams) (*fw.PagingResult, error) {
	return m.levelRepo.QueryPaging(p)
}
