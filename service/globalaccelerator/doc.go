// Code generated by smithy-go-codegen DO NOT EDIT.

// Package globalaccelerator provides the API client, operations, and parameter
// types for AWS Global Accelerator.
//
// Global Accelerator This is the Global Accelerator API Reference. This guide is
// for developers who need detailed information about Global Accelerator API
// actions, data types, and errors. For more information about Global Accelerator
// features, see the Global Accelerator Developer Guide
// (https://docs.aws.amazon.com/global-accelerator/latest/dg/what-is-global-accelerator.html).
// Global Accelerator is a service in which you create accelerators to improve the
// performance of your applications for local and global users. Depending on the
// type of accelerator you choose, you can gain additional benefits.
//
// * By using a
// standard accelerator, you can improve availability of your internet applications
// that are used by a global audience. With a standard accelerator, Global
// Accelerator directs traffic to optimal endpoints over the Amazon Web Services
// global network.
//
// * For other scenarios, you might choose a custom routing
// accelerator. With a custom routing accelerator, you can use application logic to
// directly map one or more users to a specific endpoint among many
// endpoints.
//
// Global Accelerator is a global service that supports endpoints in
// multiple Amazon Web Services Regions but you must specify the US West (Oregon)
// Region to create, update, or otherwise work with accelerators. That is, for
// example, specify --region us-west-2 on AWS CLI commands. By default, Global
// Accelerator provides you with static IP addresses that you associate with your
// accelerator. The static IP addresses are anycast from the Amazon Web Services
// edge network. For IPv4, Global Accelerator provides two static IPv4 addresses.
// For dual-stack, Global Accelerator provides a total of four addresses: two
// static IPv4 addresses and two static IPv6 addresses. With a standard accelerator
// for IPv4, instead of using the addresses that Global Accelerator provides, you
// can configure these entry points to be IPv4 addresses from your own IP address
// ranges that you bring toGlobal Accelerator (BYOIP). For a standard accelerator,
// they distribute incoming application traffic across multiple endpoint resources
// in multiple Amazon Web Services Regions , which increases the availability of
// your applications. Endpoints for standard accelerators can be Network Load
// Balancers, Application Load Balancers, Amazon EC2 instances, or Elastic IP
// addresses that are located in one Amazon Web Services Region or multiple Amazon
// Web Services Regions. For custom routing accelerators, you map traffic that
// arrives to the static IP addresses to specific Amazon EC2 servers in endpoints
// that are virtual private cloud (VPC) subnets. The static IP addresses remain
// assigned to your accelerator for as long as it exists, even if you disable the
// accelerator and it no longer accepts or routes traffic. However, when you delete
// an accelerator, you lose the static IP addresses that are assigned to it, so you
// can no longer route traffic by using them. You can use IAM policies like
// tag-based permissions with Global Accelerator to limit the users who have
// permissions to delete an accelerator. For more information, see Tag-based
// policies
// (https://docs.aws.amazon.com/global-accelerator/latest/dg/access-control-manage-access-tag-policies.html).
// For standard accelerators, Global Accelerator uses the Amazon Web Services
// global network to route traffic to the optimal regional endpoint based on
// health, client location, and policies that you configure. The service reacts
// instantly to changes in health or configuration to ensure that internet traffic
// from clients is always directed to healthy endpoints. For more information about
// understanding and using Global Accelerator, see the Global Accelerator Developer
// Guide
// (https://docs.aws.amazon.com/global-accelerator/latest/dg/what-is-global-accelerator.html).
package globalaccelerator
