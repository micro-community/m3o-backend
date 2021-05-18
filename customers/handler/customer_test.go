package handler

import (
	pb "github.com/m3o/services/customers/proto"
	mt "github.com/m3o/services/internal/test"
	"github.com/m3o/services/internal/test/fakes"
	mevents "github.com/micro/micro/v3/service/events"
	mstore "github.com/micro/micro/v3/service/store"
	"github.com/micro/micro/v3/service/store/memory"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func mockedCustomer() *Customers {
	// setting up stream and store here is a bit like resetting per test (just make sure we don't run these tests in parallel)
	mevents.DefaultStream = &fakes.FakeStream{}
	mstore.DefaultStore = memory.NewStore()
	return &Customers{
		accountsService: &fakes.FakeAccountsService{},
	}
}

func TestCreateAndDelete(t *testing.T) {
	g := NewWithT(t)
	custSvc := mockedCustomer()
	fstream := mevents.DefaultStream.(*fakes.FakeStream)
	accSvc := custSvc.accountsService.(*fakes.FakeAccountsService)
	adminCtx := mt.ContextWithAccount("micro", "foo")

	// create
	cRsp := &pb.CreateResponse{}
	err := custSvc.Create(adminCtx, &pb.CreateRequest{
		Email: "foo@bar.com",
	}, cRsp)
	g.Expect(err).To(BeNil())
	g.Expect(fstream.PublishCallCount()).To(Equal(1))

	err = custSvc.Read(adminCtx, &pb.ReadRequest{Id: cRsp.Customer.Id}, &pb.ReadResponse{})
	g.Expect(err).To(BeNil())

	// delete
	err = custSvc.Delete(adminCtx, &pb.DeleteRequest{
		Id: cRsp.Customer.Id,
	}, &pb.DeleteResponse{})
	g.Expect(err).To(BeNil())
	g.Expect(fstream.PublishCallCount()).To(Equal(2))
	g.Expect(accSvc.DeleteCallCount()).To(Equal(1))
	g.Expect(fstream.PublishCallCount()).To(Equal(2))

	rRsp := &pb.ReadResponse{}
	err = custSvc.Read(adminCtx, &pb.ReadRequest{Id: cRsp.Customer.Id}, rRsp)
	g.Expect(err).To(HaveOccurred())
}

func TestCreateAndDeleteNoOwnedNamespaces(t *testing.T) {
	g := NewWithT(t)
	custSvc := mockedCustomer()
	fstream := mevents.DefaultStream.(*fakes.FakeStream)
	accSvc := custSvc.accountsService.(*fakes.FakeAccountsService)
	adminCtx := mt.ContextWithAccount("micro", "foo")

	// create
	cRsp := &pb.CreateResponse{}
	err := custSvc.Create(adminCtx, &pb.CreateRequest{
		Email: "foo@bar.com",
	}, cRsp)
	g.Expect(err).To(BeNil())
	g.Expect(fstream.PublishCallCount()).To(Equal(1))

	err = custSvc.Read(adminCtx, &pb.ReadRequest{Id: cRsp.Customer.Id}, &pb.ReadResponse{})
	g.Expect(err).To(BeNil())

	// delete
	err = custSvc.Delete(adminCtx, &pb.DeleteRequest{
		Id: cRsp.Customer.Id,
	}, &pb.DeleteResponse{})
	g.Expect(err).To(BeNil())
	g.Expect(fstream.PublishCallCount()).To(Equal(2))
	g.Expect(accSvc.DeleteCallCount()).To(Equal(1))
	g.Expect(fstream.PublishCallCount()).To(Equal(2))

	rRsp := &pb.ReadResponse{}
	err = custSvc.Read(adminCtx, &pb.ReadRequest{Id: cRsp.Customer.Id}, rRsp)
	g.Expect(err).To(HaveOccurred())

}
