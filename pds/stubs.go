package pds

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"

	comatprototypes "github.com/bluesky-social/indigo/api/atproto"
)

func (s *Server) RegisterHandlersComAtproto(e *echo.Echo) error {
	e.POST("/xrpc/com.atproto.admin.disableAccountInvites", s.HandleComAtprotoAdminDisableAccountInvites)
	e.POST("/xrpc/com.atproto.admin.disableInviteCodes", s.HandleComAtprotoAdminDisableInviteCodes)
	e.POST("/xrpc/com.atproto.admin.enableAccountInvites", s.HandleComAtprotoAdminEnableAccountInvites)
	e.GET("/xrpc/com.atproto.admin.getAccountInfo", s.HandleComAtprotoAdminGetAccountInfo)
	e.GET("/xrpc/com.atproto.admin.getInviteCodes", s.HandleComAtprotoAdminGetInviteCodes)
	e.GET("/xrpc/com.atproto.admin.getSubjectStatus", s.HandleComAtprotoAdminGetSubjectStatus)
	e.POST("/xrpc/com.atproto.admin.sendEmail", s.HandleComAtprotoAdminSendEmail)
	e.POST("/xrpc/com.atproto.admin.updateAccountEmail", s.HandleComAtprotoAdminUpdateAccountEmail)
	e.POST("/xrpc/com.atproto.admin.updateAccountHandle", s.HandleComAtprotoAdminUpdateAccountHandle)
	e.POST("/xrpc/com.atproto.admin.updateSubjectStatus", s.HandleComAtprotoAdminUpdateSubjectStatus)
	e.GET("/xrpc/com.atproto.identity.resolveHandle", s.HandleComAtprotoIdentityResolveHandle)
	e.POST("/xrpc/com.atproto.identity.updateHandle", s.HandleComAtprotoIdentityUpdateHandle)
	e.GET("/xrpc/com.atproto.label.queryLabels", s.HandleComAtprotoLabelQueryLabels)
	e.POST("/xrpc/com.atproto.moderation.createReport", s.HandleComAtprotoModerationCreateReport)
	e.POST("/xrpc/com.atproto.repo.applyWrites", s.HandleComAtprotoRepoApplyWrites)
	e.POST("/xrpc/com.atproto.repo.createRecord", s.HandleComAtprotoRepoCreateRecord)
	e.POST("/xrpc/com.atproto.repo.deleteRecord", s.HandleComAtprotoRepoDeleteRecord)
	e.GET("/xrpc/com.atproto.repo.describeRepo", s.HandleComAtprotoRepoDescribeRepo)
	e.GET("/xrpc/com.atproto.repo.getRecord", s.HandleComAtprotoRepoGetRecord)
	e.GET("/xrpc/com.atproto.repo.listRecords", s.HandleComAtprotoRepoListRecords)
	e.POST("/xrpc/com.atproto.repo.putRecord", s.HandleComAtprotoRepoPutRecord)
	e.POST("/xrpc/com.atproto.repo.uploadBlob", s.HandleComAtprotoRepoUploadBlob)
	e.POST("/xrpc/com.atproto.server.confirmEmail", s.HandleComAtprotoServerConfirmEmail)
	e.POST("/xrpc/com.atproto.server.createAccount", s.HandleComAtprotoServerCreateAccount)
	e.POST("/xrpc/com.atproto.server.createAppPassword", s.HandleComAtprotoServerCreateAppPassword)
	e.POST("/xrpc/com.atproto.server.createInviteCode", s.HandleComAtprotoServerCreateInviteCode)
	e.POST("/xrpc/com.atproto.server.createInviteCodes", s.HandleComAtprotoServerCreateInviteCodes)
	e.POST("/xrpc/com.atproto.server.createSession", s.HandleComAtprotoServerCreateSession)
	e.POST("/xrpc/com.atproto.server.deleteAccount", s.HandleComAtprotoServerDeleteAccount)
	e.POST("/xrpc/com.atproto.server.deleteSession", s.HandleComAtprotoServerDeleteSession)
	e.GET("/xrpc/com.atproto.server.describeServer", s.HandleComAtprotoServerDescribeServer)
	e.GET("/xrpc/com.atproto.server.getAccountInviteCodes", s.HandleComAtprotoServerGetAccountInviteCodes)
	e.GET("/xrpc/com.atproto.server.getSession", s.HandleComAtprotoServerGetSession)
	e.GET("/xrpc/com.atproto.server.listAppPasswords", s.HandleComAtprotoServerListAppPasswords)
	e.POST("/xrpc/com.atproto.server.refreshSession", s.HandleComAtprotoServerRefreshSession)
	e.POST("/xrpc/com.atproto.server.requestAccountDelete", s.HandleComAtprotoServerRequestAccountDelete)
	e.POST("/xrpc/com.atproto.server.requestEmailConfirmation", s.HandleComAtprotoServerRequestEmailConfirmation)
	e.POST("/xrpc/com.atproto.server.requestEmailUpdate", s.HandleComAtprotoServerRequestEmailUpdate)
	e.POST("/xrpc/com.atproto.server.requestPasswordReset", s.HandleComAtprotoServerRequestPasswordReset)
	e.POST("/xrpc/com.atproto.server.reserveSigningKey", s.HandleComAtprotoServerReserveSigningKey)
	e.POST("/xrpc/com.atproto.server.resetPassword", s.HandleComAtprotoServerResetPassword)
	e.POST("/xrpc/com.atproto.server.revokeAppPassword", s.HandleComAtprotoServerRevokeAppPassword)
	e.POST("/xrpc/com.atproto.server.updateEmail", s.HandleComAtprotoServerUpdateEmail)
	e.GET("/xrpc/com.atproto.sync.getBlob", s.HandleComAtprotoSyncGetBlob)
	e.GET("/xrpc/com.atproto.sync.getBlocks", s.HandleComAtprotoSyncGetBlocks)
	e.GET("/xrpc/com.atproto.sync.getCheckout", s.HandleComAtprotoSyncGetCheckout)
	e.GET("/xrpc/com.atproto.sync.getHead", s.HandleComAtprotoSyncGetHead)
	e.GET("/xrpc/com.atproto.sync.getLatestCommit", s.HandleComAtprotoSyncGetLatestCommit)
	e.GET("/xrpc/com.atproto.sync.getRecord", s.HandleComAtprotoSyncGetRecord)
	e.GET("/xrpc/com.atproto.sync.getRepo", s.HandleComAtprotoSyncGetRepo)
	e.GET("/xrpc/com.atproto.sync.listBlobs", s.HandleComAtprotoSyncListBlobs)
	e.GET("/xrpc/com.atproto.sync.listRepos", s.HandleComAtprotoSyncListRepos)
	e.POST("/xrpc/com.atproto.sync.notifyOfUpdate", s.HandleComAtprotoSyncNotifyOfUpdate)
	e.POST("/xrpc/com.atproto.sync.requestCrawl", s.HandleComAtprotoSyncRequestCrawl)
	e.GET("/xrpc/com.atproto.temp.fetchLabels", s.HandleComAtprotoTempFetchLabels)
	return nil
}

func (s *Server) HandleComAtprotoAdminDisableAccountInvites(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoAdminDisableAccountInvites")
	defer span.End()

	var body comatprototypes.AdminDisableAccountInvites_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoAdminDisableAccountInvites(ctx, &body)
}

func (s *Server) HandleComAtprotoAdminDisableInviteCodes(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoAdminDisableInviteCodes")
	defer span.End()

	var body comatprototypes.AdminDisableInviteCodes_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoAdminDisableInviteCodes(ctx, &body)
}

func (s *Server) HandleComAtprotoAdminEnableAccountInvites(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoAdminEnableAccountInvites")
	defer span.End()

	var body comatprototypes.AdminEnableAccountInvites_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoAdminEnableAccountInvites(ctx, &body)
}

func (s *Server) HandleComAtprotoAdminGetAccountInfo(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoAdminGetAccountInfo")
	defer span.End()
	did := c.QueryParam("did")

	out, err := s.handleComAtprotoAdminGetAccountInfo(ctx, did)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoAdminGetInviteCodes(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoAdminGetInviteCodes")
	defer span.End()
	cursor := c.QueryParam("cursor")

	var limit int
	if p := c.QueryParam("limit"); p != "" {
		var err error
		limit, err = strconv.Atoi(p)
		if err != nil {
			return err
		}
	} else {
		limit = 100
	}
	sort := c.QueryParam("sort")

	out, err := s.handleComAtprotoAdminGetInviteCodes(ctx, cursor, limit, sort)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoAdminGetSubjectStatus(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoAdminGetSubjectStatus")
	defer span.End()

	var (
		blob = c.QueryParam("blob")
		did  = c.QueryParam("did")
		uri  = c.QueryParam("uri")
	)

	out, err := s.handleComAtprotoAdminGetSubjectStatus(ctx, blob, did, uri)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoAdminSendEmail(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoAdminSendEmail")
	defer span.End()

	var body comatprototypes.AdminSendEmail_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	out, err := s.handleComAtprotoAdminSendEmail(ctx, &body)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoAdminUpdateAccountEmail(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoAdminUpdateAccountEmail")
	defer span.End()

	var body comatprototypes.AdminUpdateAccountEmail_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoAdminUpdateAccountEmail(ctx, &body)
}

func (s *Server) HandleComAtprotoAdminUpdateAccountHandle(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoAdminUpdateAccountHandle")
	defer span.End()

	var body comatprototypes.AdminUpdateAccountHandle_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoAdminUpdateAccountHandle(ctx, &body)
}

func (s *Server) HandleComAtprotoAdminUpdateSubjectStatus(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoAdminUpdateSubjectStatus")
	defer span.End()

	var body comatprototypes.AdminUpdateSubjectStatus_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	out, err := s.handleComAtprotoAdminUpdateSubjectStatus(ctx, &body)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoIdentityResolveHandle(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoIdentityResolveHandle")
	defer span.End()
	handle := c.QueryParam("handle")

	out, err := s.handleComAtprotoIdentityResolveHandle(ctx, handle)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoIdentityUpdateHandle(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoIdentityUpdateHandle")
	defer span.End()

	var body comatprototypes.IdentityUpdateHandle_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoIdentityUpdateHandle(ctx, &body)
}

func (s *Server) HandleComAtprotoLabelQueryLabels(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoLabelQueryLabels")
	defer span.End()
	cursor := c.QueryParam("cursor")

	var limit int
	if p := c.QueryParam("limit"); p != "" {
		var err error
		limit, err = strconv.Atoi(p)
		if err != nil {
			return err
		}
	} else {
		limit = 50
	}

	var (
		sources     = c.QueryParams()["sources"]
		uriPatterns = c.QueryParams()["uriPatterns"]
	)

	out, err := s.handleComAtprotoLabelQueryLabels(ctx, cursor, limit, sources, uriPatterns)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoModerationCreateReport(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoModerationCreateReport")
	defer span.End()

	var body comatprototypes.ModerationCreateReport_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	out, err := s.handleComAtprotoModerationCreateReport(ctx, &body)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoRepoApplyWrites(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoRepoApplyWrites")
	defer span.End()

	var body comatprototypes.RepoApplyWrites_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoRepoApplyWrites(ctx, &body)
}

func (s *Server) HandleComAtprotoRepoCreateRecord(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoRepoCreateRecord")
	defer span.End()

	var body comatprototypes.RepoCreateRecord_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	out, err := s.handleComAtprotoRepoCreateRecord(ctx, &body)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoRepoDeleteRecord(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoRepoDeleteRecord")
	defer span.End()

	var body comatprototypes.RepoDeleteRecord_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoRepoDeleteRecord(ctx, &body)
}

func (s *Server) HandleComAtprotoRepoDescribeRepo(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoRepoDescribeRepo")
	defer span.End()
	repo := c.QueryParam("repo")

	out, err := s.handleComAtprotoRepoDescribeRepo(ctx, repo)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoRepoGetRecord(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoRepoGetRecord")
	defer span.End()

	var (
		cid        = c.QueryParam("cid")
		collection = c.QueryParam("collection")
		repo       = c.QueryParam("repo")
		rkey       = c.QueryParam("rkey")
	)

	out, err := s.handleComAtprotoRepoGetRecord(ctx, cid, collection, repo, rkey)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoRepoListRecords(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoRepoListRecords")
	defer span.End()
	collection := c.QueryParam("collection")
	cursor := c.QueryParam("cursor")

	var limit int
	if p := c.QueryParam("limit"); p != "" {
		var err error
		limit, err = strconv.Atoi(p)
		if err != nil {
			return err
		}
	} else {
		limit = 50
	}
	repo := c.QueryParam("repo")

	var reverse *bool
	if p := c.QueryParam("reverse"); p != "" {
		reverse_val, err := strconv.ParseBool(p)
		if err != nil {
			return err
		}
		reverse = &reverse_val
	}
	rkeyEnd := c.QueryParam("rkeyEnd")
	rkeyStart := c.QueryParam("rkeyStart")

	out, err := s.handleComAtprotoRepoListRecords(ctx, collection, cursor, limit, repo, reverse, rkeyEnd, rkeyStart)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoRepoPutRecord(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoRepoPutRecord")
	defer span.End()

	var body comatprototypes.RepoPutRecord_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	out, err := s.handleComAtprotoRepoPutRecord(ctx, &body)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoRepoUploadBlob(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoRepoUploadBlob")
	defer span.End()
	body := c.Request().Body
	contentType := c.Request().Header.Get("Content-Type")

	out, err := s.handleComAtprotoRepoUploadBlob(ctx, body, contentType)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoServerConfirmEmail(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerConfirmEmail")
	defer span.End()

	var body comatprototypes.ServerConfirmEmail_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoServerConfirmEmail(ctx, &body)
}

func (s *Server) HandleComAtprotoServerCreateAccount(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerCreateAccount")
	defer span.End()

	var body comatprototypes.ServerCreateAccount_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	out, err := s.handleComAtprotoServerCreateAccount(ctx, &body)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoServerCreateAppPassword(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerCreateAppPassword")
	defer span.End()

	var body comatprototypes.ServerCreateAppPassword_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	out, err := s.handleComAtprotoServerCreateAppPassword(ctx, &body)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoServerCreateInviteCode(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerCreateInviteCode")
	defer span.End()

	var body comatprototypes.ServerCreateInviteCode_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	out, err := s.handleComAtprotoServerCreateInviteCode(ctx, &body)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoServerCreateInviteCodes(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerCreateInviteCodes")
	defer span.End()

	var body comatprototypes.ServerCreateInviteCodes_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	out, err := s.handleComAtprotoServerCreateInviteCodes(ctx, &body)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoServerCreateSession(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerCreateSession")
	defer span.End()

	var body comatprototypes.ServerCreateSession_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	out, err := s.handleComAtprotoServerCreateSession(ctx, &body)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoServerDeleteAccount(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerDeleteAccount")
	defer span.End()

	var body comatprototypes.ServerDeleteAccount_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoServerDeleteAccount(ctx, &body)
}

func (s *Server) HandleComAtprotoServerDeleteSession(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerDeleteSession")
	defer span.End()

	return s.handleComAtprotoServerDeleteSession(ctx)
}

func (s *Server) HandleComAtprotoServerDescribeServer(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerDescribeServer")
	defer span.End()

	out, err := s.handleComAtprotoServerDescribeServer(ctx)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoServerGetAccountInviteCodes(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerGetAccountInviteCodes")
	defer span.End()

	var createAvailable bool
	if p := c.QueryParam("createAvailable"); p != "" {
		var err error
		createAvailable, err = strconv.ParseBool(p)
		if err != nil {
			return err
		}
	} else {
		createAvailable = true
	}

	var includeUsed bool
	if p := c.QueryParam("includeUsed"); p != "" {
		var err error
		includeUsed, err = strconv.ParseBool(p)
		if err != nil {
			return err
		}
	} else {
		includeUsed = true
	}

	out, err := s.handleComAtprotoServerGetAccountInviteCodes(ctx, createAvailable, includeUsed)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoServerGetSession(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerGetSession")
	defer span.End()

	out, err := s.handleComAtprotoServerGetSession(ctx)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoServerListAppPasswords(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerListAppPasswords")
	defer span.End()

	out, err := s.handleComAtprotoServerListAppPasswords(ctx)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoServerRefreshSession(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerRefreshSession")
	defer span.End()

	out, err := s.handleComAtprotoServerRefreshSession(ctx)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoServerRequestAccountDelete(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerRequestAccountDelete")
	defer span.End()

	return s.handleComAtprotoServerRequestAccountDelete(ctx)
}

func (s *Server) HandleComAtprotoServerRequestEmailConfirmation(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerRequestEmailConfirmation")
	defer span.End()

	return s.handleComAtprotoServerRequestEmailConfirmation(ctx)
}

func (s *Server) HandleComAtprotoServerRequestEmailUpdate(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerRequestEmailUpdate")
	defer span.End()

	out, err := s.handleComAtprotoServerRequestEmailUpdate(ctx)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoServerRequestPasswordReset(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerRequestPasswordReset")
	defer span.End()

	var body comatprototypes.ServerRequestPasswordReset_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoServerRequestPasswordReset(ctx, &body)
}

func (s *Server) HandleComAtprotoServerReserveSigningKey(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerReserveSigningKey")
	defer span.End()

	var body comatprototypes.ServerReserveSigningKey_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	out, err := s.handleComAtprotoServerReserveSigningKey(ctx, &body)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoServerResetPassword(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerResetPassword")
	defer span.End()

	var body comatprototypes.ServerResetPassword_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoServerResetPassword(ctx, &body)
}

func (s *Server) HandleComAtprotoServerRevokeAppPassword(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerRevokeAppPassword")
	defer span.End()

	var body comatprototypes.ServerRevokeAppPassword_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoServerRevokeAppPassword(ctx, &body)
}

func (s *Server) HandleComAtprotoServerUpdateEmail(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoServerUpdateEmail")
	defer span.End()

	var body comatprototypes.ServerUpdateEmail_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoServerUpdateEmail(ctx, &body)
}

func (s *Server) HandleComAtprotoSyncGetBlob(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncGetBlob")
	defer span.End()
	cid := c.QueryParam("cid")
	did := c.QueryParam("did")

	out, err := s.handleComAtprotoSyncGetBlob(ctx, cid, did)
	if err != nil {
		return err
	}
	return c.Stream(200, "application/octet-stream", out)
}

func (s *Server) HandleComAtprotoSyncGetBlocks(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncGetBlocks")
	defer span.End()

	var (
		cids = c.QueryParams()["cids"]
		did  = c.QueryParam("did")
	)

	out, err := s.handleComAtprotoSyncGetBlocks(ctx, cids, did)
	if err != nil {
		return err
	}
	return c.Stream(200, "application/vnd.ipld.car", out)
}

func (s *Server) HandleComAtprotoSyncGetCheckout(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncGetCheckout")
	defer span.End()
	did := c.QueryParam("did")

	out, err := s.handleComAtprotoSyncGetCheckout(ctx, did)
	if err != nil {
		return err
	}
	return c.Stream(200, "application/vnd.ipld.car", out)
}

func (s *Server) HandleComAtprotoSyncGetHead(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncGetHead")
	defer span.End()
	did := c.QueryParam("did")

	out, err := s.handleComAtprotoSyncGetHead(ctx, did)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoSyncGetLatestCommit(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncGetLatestCommit")
	defer span.End()
	did := c.QueryParam("did")

	out, err := s.handleComAtprotoSyncGetLatestCommit(ctx, did)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoSyncGetRecord(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncGetRecord")
	defer span.End()
	collection := c.QueryParam("collection")
	commit := c.QueryParam("commit")
	did := c.QueryParam("did")
	rkey := c.QueryParam("rkey")

	out, err := s.handleComAtprotoSyncGetRecord(ctx, collection, commit, did, rkey)
	if err != nil {
		return err
	}
	return c.Stream(200, "application/vnd.ipld.car", out)
}

func (s *Server) HandleComAtprotoSyncGetRepo(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncGetRepo")
	defer span.End()
	did := c.QueryParam("did")
	since := c.QueryParam("since")

	out, err := s.handleComAtprotoSyncGetRepo(ctx, did, since)
	if err != nil {
		return err
	}
	return c.Stream(200, "application/vnd.ipld.car", out)
}

func (s *Server) HandleComAtprotoSyncListBlobs(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncListBlobs")
	defer span.End()
	cursor := c.QueryParam("cursor")
	did := c.QueryParam("did")

	var limit int
	if p := c.QueryParam("limit"); p != "" {
		var err error
		limit, err = strconv.Atoi(p)
		if err != nil {
			return err
		}
	} else {
		limit = 500
	}
	since := c.QueryParam("since")

	out, err := s.handleComAtprotoSyncListBlobs(ctx, cursor, did, limit, since)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoSyncListRepos(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncListRepos")
	defer span.End()
	cursor := c.QueryParam("cursor")

	var limit int
	if p := c.QueryParam("limit"); p != "" {
		var err error
		limit, err = strconv.Atoi(p)
		if err != nil {
			return err
		}
	} else {
		limit = 500
	}

	out, err := s.handleComAtprotoSyncListRepos(ctx, cursor, limit)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s *Server) HandleComAtprotoSyncNotifyOfUpdate(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncNotifyOfUpdate")
	defer span.End()

	var body comatprototypes.SyncNotifyOfUpdate_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoSyncNotifyOfUpdate(ctx, &body)
}

func (s *Server) HandleComAtprotoSyncRequestCrawl(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncRequestCrawl")
	defer span.End()

	var body comatprototypes.SyncRequestCrawl_Input
	if err := c.Bind(&body); err != nil {
		return err
	}

	return s.handleComAtprotoSyncRequestCrawl(ctx, &body)
}

func (s *Server) HandleComAtprotoTempFetchLabels(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoTempFetchLabels")
	defer span.End()

	var limit int
	if p := c.QueryParam("limit"); p != "" {
		var err error
		limit, err = strconv.Atoi(p)
		if err != nil {
			return err
		}
	} else {
		limit = 50
	}

	var since *int
	if p := c.QueryParam("since"); p != "" {
		sinceVal, err := strconv.Atoi(p)
		if err != nil {
			return err
		}
		since = &sinceVal
	}

	out, err := s.handleComAtprotoTempFetchLabels(ctx, limit, since)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}
