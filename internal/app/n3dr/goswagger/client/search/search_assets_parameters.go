// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSearchAssetsParams creates a new SearchAssetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchAssetsParams() *SearchAssetsParams {
	return &SearchAssetsParams{
		requestTimeout: cr.DefaultTimeout,
	}
}

// NewSearchAssetsParamsWithTimeout creates a new SearchAssetsParams object
// with the ability to set a timeout on a request.
func NewSearchAssetsParamsWithTimeout(timeout time.Duration) *SearchAssetsParams {
	return &SearchAssetsParams{
		requestTimeout: timeout,
	}
}

// NewSearchAssetsParamsWithContext creates a new SearchAssetsParams object
// with the ability to set a context for a request.
func NewSearchAssetsParamsWithContext(ctx context.Context) *SearchAssetsParams {
	return &SearchAssetsParams{
		Context: ctx,
	}
}

// NewSearchAssetsParamsWithHTTPClient creates a new SearchAssetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchAssetsParamsWithHTTPClient(client *http.Client) *SearchAssetsParams {
	return &SearchAssetsParams{
		HTTPClient: client,
	}
}

/* SearchAssetsParams contains all the parameters to send to the API endpoint
   for the search assets operation.

   Typically these are written to a http.Request.
*/
type SearchAssetsParams struct {

	/* ConanBaseVersion.

	   Conan base version
	*/
	ConanBaseVersion *string

	/* ConanChannel.

	   Conan channel
	*/
	ConanChannel *string

	/* ContinuationToken.

	   A token returned by a prior request. If present, the next page of results are returned
	*/
	ContinuationToken *string

	/* Direction.

	   The direction to sort records in, defaults to ascending ('asc') for all sort fields, except version, which defaults to descending ('desc')
	*/
	Direction string

	/* DockerContentDigest.

	   Docker content digest
	*/
	DockerContentDigest *string

	/* DockerImageName.

	   Docker image name
	*/
	DockerImageName *string

	/* DockerImageTag.

	   Docker image tag
	*/
	DockerImageTag *string

	/* DockerLayerID.

	   Docker layer ID
	*/
	DockerLayerID *string

	/* Format.

	   Query by format
	*/
	Format *string

	/* Group.

	   Component group
	*/
	Group *string

	/* MavenArtifactID.

	   Maven artifactId
	*/
	MavenArtifactID *string

	/* MavenBaseVersion.

	   Maven base version
	*/
	MavenBaseVersion *string

	/* MavenClassifier.

	   Maven classifier of component's asset
	*/
	MavenClassifier *string

	/* MavenExtension.

	   Maven extension of component's asset
	*/
	MavenExtension *string

	/* MavenGroupID.

	   Maven groupId
	*/
	MavenGroupID *string

	/* Md5.

	   Specific MD5 hash of component's asset
	*/
	Md5 *string

	/* Name.

	   Component name
	*/
	Name *string

	/* NpmScope.

	   npm scope
	*/
	NpmScope *string

	/* NugetID.

	   NuGet id
	*/
	NugetID *string

	/* NugetTags.

	   NuGet tags
	*/
	NugetTags *string

	/* P2PluginName.

	   p2 plugin name
	*/
	P2PluginName *string

	/* Prerelease.

	   Prerelease version flag
	*/
	Prerelease *string

	/* PypiClassifiers.

	   PyPI classifiers
	*/
	PypiClassifiers *string

	/* PypiDescription.

	   PyPI description
	*/
	PypiDescription *string

	/* PypiKeywords.

	   PyPI keywords
	*/
	PypiKeywords *string

	/* PypiSummary.

	   PyPI summary
	*/
	PypiSummary *string

	/* Q.

	   Query by keyword
	*/
	Q *string

	/* Repository.

	   Repository name
	*/
	Repository *string

	/* RubygemsDescription.

	   RubyGems description
	*/
	RubygemsDescription *string

	/* RubygemsPlatform.

	   RubyGems platform
	*/
	RubygemsPlatform *string

	/* RubygemsSummary.

	   RubyGems summary
	*/
	RubygemsSummary *string

	/* Sha1.

	   Specific SHA-1 hash of component's asset
	*/
	Sha1 *string

	/* Sha256.

	   Specific SHA-256 hash of component's asset
	*/
	Sha256 *string

	/* Sha512.

	   Specific SHA-512 hash of component's asset
	*/
	Sha512 *string

	/* Sort.

	   The field to sort the results against, if left empty, a sort based on match weight will be used.
	*/
	Sort string

	/* Timeout.

	   How long to wait for search results in seconds. If this value is not provided, the system default timeout will be used.

	   Format: int32
	*/
	Timeout int32

	/* Version.

	   Component version
	*/
	Version *string

	/* YumArchitecture.

	   Yum architecture
	*/
	YumArchitecture *string

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithDefaults hydrates default values in the search assets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchAssetsParams) WithDefaults() *SearchAssetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search assets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchAssetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithRequestTimeout adds the timeout to the search assets params
func (o *SearchAssetsParams) WithRequestTimeout(timeout time.Duration) *SearchAssetsParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the search assets params
func (o *SearchAssetsParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the search assets params
func (o *SearchAssetsParams) WithContext(ctx context.Context) *SearchAssetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search assets params
func (o *SearchAssetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search assets params
func (o *SearchAssetsParams) WithHTTPClient(client *http.Client) *SearchAssetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search assets params
func (o *SearchAssetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConanBaseVersion adds the conanBaseVersion to the search assets params
func (o *SearchAssetsParams) WithConanBaseVersion(conanBaseVersion *string) *SearchAssetsParams {
	o.SetConanBaseVersion(conanBaseVersion)
	return o
}

// SetConanBaseVersion adds the conanBaseVersion to the search assets params
func (o *SearchAssetsParams) SetConanBaseVersion(conanBaseVersion *string) {
	o.ConanBaseVersion = conanBaseVersion
}

// WithConanChannel adds the conanChannel to the search assets params
func (o *SearchAssetsParams) WithConanChannel(conanChannel *string) *SearchAssetsParams {
	o.SetConanChannel(conanChannel)
	return o
}

// SetConanChannel adds the conanChannel to the search assets params
func (o *SearchAssetsParams) SetConanChannel(conanChannel *string) {
	o.ConanChannel = conanChannel
}

// WithContinuationToken adds the continuationToken to the search assets params
func (o *SearchAssetsParams) WithContinuationToken(continuationToken *string) *SearchAssetsParams {
	o.SetContinuationToken(continuationToken)
	return o
}

// SetContinuationToken adds the continuationToken to the search assets params
func (o *SearchAssetsParams) SetContinuationToken(continuationToken *string) {
	o.ContinuationToken = continuationToken
}

// WithDirection adds the direction to the search assets params
func (o *SearchAssetsParams) WithDirection(direction string) *SearchAssetsParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the search assets params
func (o *SearchAssetsParams) SetDirection(direction string) {
	o.Direction = direction
}

// WithDockerContentDigest adds the dockerContentDigest to the search assets params
func (o *SearchAssetsParams) WithDockerContentDigest(dockerContentDigest *string) *SearchAssetsParams {
	o.SetDockerContentDigest(dockerContentDigest)
	return o
}

// SetDockerContentDigest adds the dockerContentDigest to the search assets params
func (o *SearchAssetsParams) SetDockerContentDigest(dockerContentDigest *string) {
	o.DockerContentDigest = dockerContentDigest
}

// WithDockerImageName adds the dockerImageName to the search assets params
func (o *SearchAssetsParams) WithDockerImageName(dockerImageName *string) *SearchAssetsParams {
	o.SetDockerImageName(dockerImageName)
	return o
}

// SetDockerImageName adds the dockerImageName to the search assets params
func (o *SearchAssetsParams) SetDockerImageName(dockerImageName *string) {
	o.DockerImageName = dockerImageName
}

// WithDockerImageTag adds the dockerImageTag to the search assets params
func (o *SearchAssetsParams) WithDockerImageTag(dockerImageTag *string) *SearchAssetsParams {
	o.SetDockerImageTag(dockerImageTag)
	return o
}

// SetDockerImageTag adds the dockerImageTag to the search assets params
func (o *SearchAssetsParams) SetDockerImageTag(dockerImageTag *string) {
	o.DockerImageTag = dockerImageTag
}

// WithDockerLayerID adds the dockerLayerID to the search assets params
func (o *SearchAssetsParams) WithDockerLayerID(dockerLayerID *string) *SearchAssetsParams {
	o.SetDockerLayerID(dockerLayerID)
	return o
}

// SetDockerLayerID adds the dockerLayerId to the search assets params
func (o *SearchAssetsParams) SetDockerLayerID(dockerLayerID *string) {
	o.DockerLayerID = dockerLayerID
}

// WithFormat adds the format to the search assets params
func (o *SearchAssetsParams) WithFormat(format *string) *SearchAssetsParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the search assets params
func (o *SearchAssetsParams) SetFormat(format *string) {
	o.Format = format
}

// WithGroup adds the group to the search assets params
func (o *SearchAssetsParams) WithGroup(group *string) *SearchAssetsParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the search assets params
func (o *SearchAssetsParams) SetGroup(group *string) {
	o.Group = group
}

// WithMavenArtifactID adds the mavenArtifactID to the search assets params
func (o *SearchAssetsParams) WithMavenArtifactID(mavenArtifactID *string) *SearchAssetsParams {
	o.SetMavenArtifactID(mavenArtifactID)
	return o
}

// SetMavenArtifactID adds the mavenArtifactId to the search assets params
func (o *SearchAssetsParams) SetMavenArtifactID(mavenArtifactID *string) {
	o.MavenArtifactID = mavenArtifactID
}

// WithMavenBaseVersion adds the mavenBaseVersion to the search assets params
func (o *SearchAssetsParams) WithMavenBaseVersion(mavenBaseVersion *string) *SearchAssetsParams {
	o.SetMavenBaseVersion(mavenBaseVersion)
	return o
}

// SetMavenBaseVersion adds the mavenBaseVersion to the search assets params
func (o *SearchAssetsParams) SetMavenBaseVersion(mavenBaseVersion *string) {
	o.MavenBaseVersion = mavenBaseVersion
}

// WithMavenClassifier adds the mavenClassifier to the search assets params
func (o *SearchAssetsParams) WithMavenClassifier(mavenClassifier *string) *SearchAssetsParams {
	o.SetMavenClassifier(mavenClassifier)
	return o
}

// SetMavenClassifier adds the mavenClassifier to the search assets params
func (o *SearchAssetsParams) SetMavenClassifier(mavenClassifier *string) {
	o.MavenClassifier = mavenClassifier
}

// WithMavenExtension adds the mavenExtension to the search assets params
func (o *SearchAssetsParams) WithMavenExtension(mavenExtension *string) *SearchAssetsParams {
	o.SetMavenExtension(mavenExtension)
	return o
}

// SetMavenExtension adds the mavenExtension to the search assets params
func (o *SearchAssetsParams) SetMavenExtension(mavenExtension *string) {
	o.MavenExtension = mavenExtension
}

// WithMavenGroupID adds the mavenGroupID to the search assets params
func (o *SearchAssetsParams) WithMavenGroupID(mavenGroupID *string) *SearchAssetsParams {
	o.SetMavenGroupID(mavenGroupID)
	return o
}

// SetMavenGroupID adds the mavenGroupId to the search assets params
func (o *SearchAssetsParams) SetMavenGroupID(mavenGroupID *string) {
	o.MavenGroupID = mavenGroupID
}

// WithMd5 adds the md5 to the search assets params
func (o *SearchAssetsParams) WithMd5(md5 *string) *SearchAssetsParams {
	o.SetMd5(md5)
	return o
}

// SetMd5 adds the md5 to the search assets params
func (o *SearchAssetsParams) SetMd5(md5 *string) {
	o.Md5 = md5
}

// WithName adds the name to the search assets params
func (o *SearchAssetsParams) WithName(name *string) *SearchAssetsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the search assets params
func (o *SearchAssetsParams) SetName(name *string) {
	o.Name = name
}

// WithNpmScope adds the npmScope to the search assets params
func (o *SearchAssetsParams) WithNpmScope(npmScope *string) *SearchAssetsParams {
	o.SetNpmScope(npmScope)
	return o
}

// SetNpmScope adds the npmScope to the search assets params
func (o *SearchAssetsParams) SetNpmScope(npmScope *string) {
	o.NpmScope = npmScope
}

// WithNugetID adds the nugetID to the search assets params
func (o *SearchAssetsParams) WithNugetID(nugetID *string) *SearchAssetsParams {
	o.SetNugetID(nugetID)
	return o
}

// SetNugetID adds the nugetId to the search assets params
func (o *SearchAssetsParams) SetNugetID(nugetID *string) {
	o.NugetID = nugetID
}

// WithNugetTags adds the nugetTags to the search assets params
func (o *SearchAssetsParams) WithNugetTags(nugetTags *string) *SearchAssetsParams {
	o.SetNugetTags(nugetTags)
	return o
}

// SetNugetTags adds the nugetTags to the search assets params
func (o *SearchAssetsParams) SetNugetTags(nugetTags *string) {
	o.NugetTags = nugetTags
}

// WithP2PluginName adds the p2PluginName to the search assets params
func (o *SearchAssetsParams) WithP2PluginName(p2PluginName *string) *SearchAssetsParams {
	o.SetP2PluginName(p2PluginName)
	return o
}

// SetP2PluginName adds the p2PluginName to the search assets params
func (o *SearchAssetsParams) SetP2PluginName(p2PluginName *string) {
	o.P2PluginName = p2PluginName
}

// WithPrerelease adds the prerelease to the search assets params
func (o *SearchAssetsParams) WithPrerelease(prerelease *string) *SearchAssetsParams {
	o.SetPrerelease(prerelease)
	return o
}

// SetPrerelease adds the prerelease to the search assets params
func (o *SearchAssetsParams) SetPrerelease(prerelease *string) {
	o.Prerelease = prerelease
}

// WithPypiClassifiers adds the pypiClassifiers to the search assets params
func (o *SearchAssetsParams) WithPypiClassifiers(pypiClassifiers *string) *SearchAssetsParams {
	o.SetPypiClassifiers(pypiClassifiers)
	return o
}

// SetPypiClassifiers adds the pypiClassifiers to the search assets params
func (o *SearchAssetsParams) SetPypiClassifiers(pypiClassifiers *string) {
	o.PypiClassifiers = pypiClassifiers
}

// WithPypiDescription adds the pypiDescription to the search assets params
func (o *SearchAssetsParams) WithPypiDescription(pypiDescription *string) *SearchAssetsParams {
	o.SetPypiDescription(pypiDescription)
	return o
}

// SetPypiDescription adds the pypiDescription to the search assets params
func (o *SearchAssetsParams) SetPypiDescription(pypiDescription *string) {
	o.PypiDescription = pypiDescription
}

// WithPypiKeywords adds the pypiKeywords to the search assets params
func (o *SearchAssetsParams) WithPypiKeywords(pypiKeywords *string) *SearchAssetsParams {
	o.SetPypiKeywords(pypiKeywords)
	return o
}

// SetPypiKeywords adds the pypiKeywords to the search assets params
func (o *SearchAssetsParams) SetPypiKeywords(pypiKeywords *string) {
	o.PypiKeywords = pypiKeywords
}

// WithPypiSummary adds the pypiSummary to the search assets params
func (o *SearchAssetsParams) WithPypiSummary(pypiSummary *string) *SearchAssetsParams {
	o.SetPypiSummary(pypiSummary)
	return o
}

// SetPypiSummary adds the pypiSummary to the search assets params
func (o *SearchAssetsParams) SetPypiSummary(pypiSummary *string) {
	o.PypiSummary = pypiSummary
}

// WithQ adds the q to the search assets params
func (o *SearchAssetsParams) WithQ(q *string) *SearchAssetsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the search assets params
func (o *SearchAssetsParams) SetQ(q *string) {
	o.Q = q
}

// WithRepository adds the repository to the search assets params
func (o *SearchAssetsParams) WithRepository(repository *string) *SearchAssetsParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the search assets params
func (o *SearchAssetsParams) SetRepository(repository *string) {
	o.Repository = repository
}

// WithRubygemsDescription adds the rubygemsDescription to the search assets params
func (o *SearchAssetsParams) WithRubygemsDescription(rubygemsDescription *string) *SearchAssetsParams {
	o.SetRubygemsDescription(rubygemsDescription)
	return o
}

// SetRubygemsDescription adds the rubygemsDescription to the search assets params
func (o *SearchAssetsParams) SetRubygemsDescription(rubygemsDescription *string) {
	o.RubygemsDescription = rubygemsDescription
}

// WithRubygemsPlatform adds the rubygemsPlatform to the search assets params
func (o *SearchAssetsParams) WithRubygemsPlatform(rubygemsPlatform *string) *SearchAssetsParams {
	o.SetRubygemsPlatform(rubygemsPlatform)
	return o
}

// SetRubygemsPlatform adds the rubygemsPlatform to the search assets params
func (o *SearchAssetsParams) SetRubygemsPlatform(rubygemsPlatform *string) {
	o.RubygemsPlatform = rubygemsPlatform
}

// WithRubygemsSummary adds the rubygemsSummary to the search assets params
func (o *SearchAssetsParams) WithRubygemsSummary(rubygemsSummary *string) *SearchAssetsParams {
	o.SetRubygemsSummary(rubygemsSummary)
	return o
}

// SetRubygemsSummary adds the rubygemsSummary to the search assets params
func (o *SearchAssetsParams) SetRubygemsSummary(rubygemsSummary *string) {
	o.RubygemsSummary = rubygemsSummary
}

// WithSha1 adds the sha1 to the search assets params
func (o *SearchAssetsParams) WithSha1(sha1 *string) *SearchAssetsParams {
	o.SetSha1(sha1)
	return o
}

// SetSha1 adds the sha1 to the search assets params
func (o *SearchAssetsParams) SetSha1(sha1 *string) {
	o.Sha1 = sha1
}

// WithSha256 adds the sha256 to the search assets params
func (o *SearchAssetsParams) WithSha256(sha256 *string) *SearchAssetsParams {
	o.SetSha256(sha256)
	return o
}

// SetSha256 adds the sha256 to the search assets params
func (o *SearchAssetsParams) SetSha256(sha256 *string) {
	o.Sha256 = sha256
}

// WithSha512 adds the sha512 to the search assets params
func (o *SearchAssetsParams) WithSha512(sha512 *string) *SearchAssetsParams {
	o.SetSha512(sha512)
	return o
}

// SetSha512 adds the sha512 to the search assets params
func (o *SearchAssetsParams) SetSha512(sha512 *string) {
	o.Sha512 = sha512
}

// WithSort adds the sort to the search assets params
func (o *SearchAssetsParams) WithSort(sort string) *SearchAssetsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search assets params
func (o *SearchAssetsParams) SetSort(sort string) {
	o.Sort = sort
}

// WithTimeout adds the timeout to the search assets params
func (o *SearchAssetsParams) WithTimeout(timeout int32) *SearchAssetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search assets params
func (o *SearchAssetsParams) SetTimeout(timeout int32) {
	o.Timeout = timeout
}

// WithVersion adds the version to the search assets params
func (o *SearchAssetsParams) WithVersion(version *string) *SearchAssetsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the search assets params
func (o *SearchAssetsParams) SetVersion(version *string) {
	o.Version = version
}

// WithYumArchitecture adds the yumArchitecture to the search assets params
func (o *SearchAssetsParams) WithYumArchitecture(yumArchitecture *string) *SearchAssetsParams {
	o.SetYumArchitecture(yumArchitecture)
	return o
}

// SetYumArchitecture adds the yumArchitecture to the search assets params
func (o *SearchAssetsParams) SetYumArchitecture(yumArchitecture *string) {
	o.YumArchitecture = yumArchitecture
}

// WriteToRequest writes these params to a swagger request
func (o *SearchAssetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	if o.ConanBaseVersion != nil {

		// query param conan.baseVersion
		var qrConanBaseVersion string

		if o.ConanBaseVersion != nil {
			qrConanBaseVersion = *o.ConanBaseVersion
		}
		qConanBaseVersion := qrConanBaseVersion
		if qConanBaseVersion != "" {

			if err := r.SetQueryParam("conan.baseVersion", qConanBaseVersion); err != nil {
				return err
			}
		}
	}

	if o.ConanChannel != nil {

		// query param conan.channel
		var qrConanChannel string

		if o.ConanChannel != nil {
			qrConanChannel = *o.ConanChannel
		}
		qConanChannel := qrConanChannel
		if qConanChannel != "" {

			if err := r.SetQueryParam("conan.channel", qConanChannel); err != nil {
				return err
			}
		}
	}

	if o.ContinuationToken != nil {

		// query param continuationToken
		var qrContinuationToken string

		if o.ContinuationToken != nil {
			qrContinuationToken = *o.ContinuationToken
		}
		qContinuationToken := qrContinuationToken
		if qContinuationToken != "" {

			if err := r.SetQueryParam("continuationToken", qContinuationToken); err != nil {
				return err
			}
		}
	}

	// query param direction
	qrDirection := o.Direction
	qDirection := qrDirection

	if err := r.SetQueryParam("direction", qDirection); err != nil {
		return err
	}

	if o.DockerContentDigest != nil {

		// query param docker.contentDigest
		var qrDockerContentDigest string

		if o.DockerContentDigest != nil {
			qrDockerContentDigest = *o.DockerContentDigest
		}
		qDockerContentDigest := qrDockerContentDigest
		if qDockerContentDigest != "" {

			if err := r.SetQueryParam("docker.contentDigest", qDockerContentDigest); err != nil {
				return err
			}
		}
	}

	if o.DockerImageName != nil {

		// query param docker.imageName
		var qrDockerImageName string

		if o.DockerImageName != nil {
			qrDockerImageName = *o.DockerImageName
		}
		qDockerImageName := qrDockerImageName
		if qDockerImageName != "" {

			if err := r.SetQueryParam("docker.imageName", qDockerImageName); err != nil {
				return err
			}
		}
	}

	if o.DockerImageTag != nil {

		// query param docker.imageTag
		var qrDockerImageTag string

		if o.DockerImageTag != nil {
			qrDockerImageTag = *o.DockerImageTag
		}
		qDockerImageTag := qrDockerImageTag
		if qDockerImageTag != "" {

			if err := r.SetQueryParam("docker.imageTag", qDockerImageTag); err != nil {
				return err
			}
		}
	}

	if o.DockerLayerID != nil {

		// query param docker.layerId
		var qrDockerLayerID string

		if o.DockerLayerID != nil {
			qrDockerLayerID = *o.DockerLayerID
		}
		qDockerLayerID := qrDockerLayerID
		if qDockerLayerID != "" {

			if err := r.SetQueryParam("docker.layerId", qDockerLayerID); err != nil {
				return err
			}
		}
	}

	if o.Format != nil {

		// query param format
		var qrFormat string

		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {

			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}
	}

	if o.Group != nil {

		// query param group
		var qrGroup string

		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := qrGroup
		if qGroup != "" {

			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}
	}

	if o.MavenArtifactID != nil {

		// query param maven.artifactId
		var qrMavenArtifactID string

		if o.MavenArtifactID != nil {
			qrMavenArtifactID = *o.MavenArtifactID
		}
		qMavenArtifactID := qrMavenArtifactID
		if qMavenArtifactID != "" {

			if err := r.SetQueryParam("maven.artifactId", qMavenArtifactID); err != nil {
				return err
			}
		}
	}

	if o.MavenBaseVersion != nil {

		// query param maven.baseVersion
		var qrMavenBaseVersion string

		if o.MavenBaseVersion != nil {
			qrMavenBaseVersion = *o.MavenBaseVersion
		}
		qMavenBaseVersion := qrMavenBaseVersion
		if qMavenBaseVersion != "" {

			if err := r.SetQueryParam("maven.baseVersion", qMavenBaseVersion); err != nil {
				return err
			}
		}
	}

	if o.MavenClassifier != nil {

		// query param maven.classifier
		var qrMavenClassifier string

		if o.MavenClassifier != nil {
			qrMavenClassifier = *o.MavenClassifier
		}
		qMavenClassifier := qrMavenClassifier
		if qMavenClassifier != "" {

			if err := r.SetQueryParam("maven.classifier", qMavenClassifier); err != nil {
				return err
			}
		}
	}

	if o.MavenExtension != nil {

		// query param maven.extension
		var qrMavenExtension string

		if o.MavenExtension != nil {
			qrMavenExtension = *o.MavenExtension
		}
		qMavenExtension := qrMavenExtension
		if qMavenExtension != "" {

			if err := r.SetQueryParam("maven.extension", qMavenExtension); err != nil {
				return err
			}
		}
	}

	if o.MavenGroupID != nil {

		// query param maven.groupId
		var qrMavenGroupID string

		if o.MavenGroupID != nil {
			qrMavenGroupID = *o.MavenGroupID
		}
		qMavenGroupID := qrMavenGroupID
		if qMavenGroupID != "" {

			if err := r.SetQueryParam("maven.groupId", qMavenGroupID); err != nil {
				return err
			}
		}
	}

	if o.Md5 != nil {

		// query param md5
		var qrMd5 string

		if o.Md5 != nil {
			qrMd5 = *o.Md5
		}
		qMd5 := qrMd5
		if qMd5 != "" {

			if err := r.SetQueryParam("md5", qMd5); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NpmScope != nil {

		// query param npm.scope
		var qrNpmScope string

		if o.NpmScope != nil {
			qrNpmScope = *o.NpmScope
		}
		qNpmScope := qrNpmScope
		if qNpmScope != "" {

			if err := r.SetQueryParam("npm.scope", qNpmScope); err != nil {
				return err
			}
		}
	}

	if o.NugetID != nil {

		// query param nuget.id
		var qrNugetID string

		if o.NugetID != nil {
			qrNugetID = *o.NugetID
		}
		qNugetID := qrNugetID
		if qNugetID != "" {

			if err := r.SetQueryParam("nuget.id", qNugetID); err != nil {
				return err
			}
		}
	}

	if o.NugetTags != nil {

		// query param nuget.tags
		var qrNugetTags string

		if o.NugetTags != nil {
			qrNugetTags = *o.NugetTags
		}
		qNugetTags := qrNugetTags
		if qNugetTags != "" {

			if err := r.SetQueryParam("nuget.tags", qNugetTags); err != nil {
				return err
			}
		}
	}

	if o.P2PluginName != nil {

		// query param p2.pluginName
		var qrP2PluginName string

		if o.P2PluginName != nil {
			qrP2PluginName = *o.P2PluginName
		}
		qP2PluginName := qrP2PluginName
		if qP2PluginName != "" {

			if err := r.SetQueryParam("p2.pluginName", qP2PluginName); err != nil {
				return err
			}
		}
	}

	if o.Prerelease != nil {

		// query param prerelease
		var qrPrerelease string

		if o.Prerelease != nil {
			qrPrerelease = *o.Prerelease
		}
		qPrerelease := qrPrerelease
		if qPrerelease != "" {

			if err := r.SetQueryParam("prerelease", qPrerelease); err != nil {
				return err
			}
		}
	}

	if o.PypiClassifiers != nil {

		// query param pypi.classifiers
		var qrPypiClassifiers string

		if o.PypiClassifiers != nil {
			qrPypiClassifiers = *o.PypiClassifiers
		}
		qPypiClassifiers := qrPypiClassifiers
		if qPypiClassifiers != "" {

			if err := r.SetQueryParam("pypi.classifiers", qPypiClassifiers); err != nil {
				return err
			}
		}
	}

	if o.PypiDescription != nil {

		// query param pypi.description
		var qrPypiDescription string

		if o.PypiDescription != nil {
			qrPypiDescription = *o.PypiDescription
		}
		qPypiDescription := qrPypiDescription
		if qPypiDescription != "" {

			if err := r.SetQueryParam("pypi.description", qPypiDescription); err != nil {
				return err
			}
		}
	}

	if o.PypiKeywords != nil {

		// query param pypi.keywords
		var qrPypiKeywords string

		if o.PypiKeywords != nil {
			qrPypiKeywords = *o.PypiKeywords
		}
		qPypiKeywords := qrPypiKeywords
		if qPypiKeywords != "" {

			if err := r.SetQueryParam("pypi.keywords", qPypiKeywords); err != nil {
				return err
			}
		}
	}

	if o.PypiSummary != nil {

		// query param pypi.summary
		var qrPypiSummary string

		if o.PypiSummary != nil {
			qrPypiSummary = *o.PypiSummary
		}
		qPypiSummary := qrPypiSummary
		if qPypiSummary != "" {

			if err := r.SetQueryParam("pypi.summary", qPypiSummary); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Repository != nil {

		// query param repository
		var qrRepository string

		if o.Repository != nil {
			qrRepository = *o.Repository
		}
		qRepository := qrRepository
		if qRepository != "" {

			if err := r.SetQueryParam("repository", qRepository); err != nil {
				return err
			}
		}
	}

	if o.RubygemsDescription != nil {

		// query param rubygems.description
		var qrRubygemsDescription string

		if o.RubygemsDescription != nil {
			qrRubygemsDescription = *o.RubygemsDescription
		}
		qRubygemsDescription := qrRubygemsDescription
		if qRubygemsDescription != "" {

			if err := r.SetQueryParam("rubygems.description", qRubygemsDescription); err != nil {
				return err
			}
		}
	}

	if o.RubygemsPlatform != nil {

		// query param rubygems.platform
		var qrRubygemsPlatform string

		if o.RubygemsPlatform != nil {
			qrRubygemsPlatform = *o.RubygemsPlatform
		}
		qRubygemsPlatform := qrRubygemsPlatform
		if qRubygemsPlatform != "" {

			if err := r.SetQueryParam("rubygems.platform", qRubygemsPlatform); err != nil {
				return err
			}
		}
	}

	if o.RubygemsSummary != nil {

		// query param rubygems.summary
		var qrRubygemsSummary string

		if o.RubygemsSummary != nil {
			qrRubygemsSummary = *o.RubygemsSummary
		}
		qRubygemsSummary := qrRubygemsSummary
		if qRubygemsSummary != "" {

			if err := r.SetQueryParam("rubygems.summary", qRubygemsSummary); err != nil {
				return err
			}
		}
	}

	if o.Sha1 != nil {

		// query param sha1
		var qrSha1 string

		if o.Sha1 != nil {
			qrSha1 = *o.Sha1
		}
		qSha1 := qrSha1
		if qSha1 != "" {

			if err := r.SetQueryParam("sha1", qSha1); err != nil {
				return err
			}
		}
	}

	if o.Sha256 != nil {

		// query param sha256
		var qrSha256 string

		if o.Sha256 != nil {
			qrSha256 = *o.Sha256
		}
		qSha256 := qrSha256
		if qSha256 != "" {

			if err := r.SetQueryParam("sha256", qSha256); err != nil {
				return err
			}
		}
	}

	if o.Sha512 != nil {

		// query param sha512
		var qrSha512 string

		if o.Sha512 != nil {
			qrSha512 = *o.Sha512
		}
		qSha512 := qrSha512
		if qSha512 != "" {

			if err := r.SetQueryParam("sha512", qSha512); err != nil {
				return err
			}
		}
	}

	// query param sort
	qrSort := o.Sort
	qSort := qrSort

	if err := r.SetQueryParam("sort", qSort); err != nil {
		return err
	}

	// query param timeout
	qrTimeout := o.Timeout
	qTimeout := swag.FormatInt32(qrTimeout)

	if err := r.SetQueryParam("timeout", qTimeout); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if o.YumArchitecture != nil {

		// query param yum.architecture
		var qrYumArchitecture string

		if o.YumArchitecture != nil {
			qrYumArchitecture = *o.YumArchitecture
		}
		qYumArchitecture := qrYumArchitecture
		if qYumArchitecture != "" {

			if err := r.SetQueryParam("yum.architecture", qYumArchitecture); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
