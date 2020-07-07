
// This file is an automatically generated Java binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package geth;

import org.ethereum.geth.*;
import java.util.*;



public class VoidTest {
	// ABI is the input ABI used to generate the binding from.
	public final static String ABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"method\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]";
	
		// VoidTestFuncSigs maps the 4-byte function signature to its string representation.
		public final static Map<String, String> VoidTestFuncSigs;
		static {
			Hashtable<String, String> temp = new Hashtable<String, String>();
			temp.put("9a40e3f6", "method(uint256)");
			
			VoidTestFuncSigs = Collections.unmodifiableMap(temp);
		}
	
	
	// BYTECODE is the compiled bytecode used for deploying new contracts.
	public final static String BYTECODE = "0x6080604052348015600f57600080fd5b50608e8061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80639a40e3f614602d575b600080fd5b603c60383660046041565b603e565b005b50565b6000602082840312156051578081fd5b503591905056fea26469706673582212206226e2422e892cda4f6d3c44a662bb7e410db375246c6b0a80341e6e7377e14f64736f6c63430006050033";

	// deploy deploys a new Ethereum contract, binding an instance of VoidTest to it.
	public static VoidTest deploy(TransactOpts auth, EthereumClient client) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		String bytecode = BYTECODE;
		
		
		return new VoidTest(Geth.deployContract(auth, ABI, Geth.decodeFromHex(bytecode), client, args));
	}

	// Internal constructor used by contract deployment.
	private VoidTest(BoundContract deployment) {
		this.Address  = deployment.getAddress();
		this.Deployer = deployment.getDeployer();
		this.Contract = deployment;
	}
	

	// Ethereum address where this contract is located at.
	public final Address Address;

	// Ethereum transaction in which this contract was deployed (if known!).
	public final Transaction Deployer;

	// Contract instance bound to a blockchain address.
	private final BoundContract Contract;

	// Creates a new instance of VoidTest, bound to a specific deployed contract.
	public VoidTest(Address address, EthereumClient client) throws Exception {
		this(Geth.bindContract(address, ABI, client));
	}

	
	

	// method is a free data retrieval call binding the contract method 0x9a40e3f6.
	//
	// Solidity: function method(uint256 a) view returns()
	public void method(CallOpts opts, BigInt a) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setBigInt(a);args.set(0,arg0);
		

		Interfaces results = Geth.newInterfaces(0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "method", args);
		
		
	}
	

	

    

    
}

